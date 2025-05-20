package final

import (
	"context"
	"errors"
	"final/interceptors"
	"final/internal/config"
	"final/internal/security/jwt"
	"final/internal/service"
	"final/internal/utils/observability/log"
	pb "final/pkg/proto/sync/final-boss/v1"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"time"
)

type Server struct {
	svc service.Service
	pb.UnimplementedAuthServer
	cfg          *config.Config
	logger       *log.Logger
	grpcServer   *grpc.Server
	grpcListener net.Listener
	httpListener *http.Server
}

func NewServer(cfg *config.Config, logger *log.Logger, svc service.Service) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		svc:    svc,
	}
}

func (s *Server) Listen() error {
	jwt.Init(s.cfg.JWT.Secret, time.Hour*24) // инициализируем глобальные параметры

	s.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.WithValidation(s.logger),
			interceptors.JwtInterceptor(s.logger, true),
		),
	)
	pb.RegisterAuthServer(s.grpcServer, s)
	reflection.Register(s.grpcServer)

	// --- HTTP Gateway (grpc-gateway) ---
	mux := http.NewServeMux()
	gwMux := runtime.NewServeMux()

	// Регистрируем grpc-gateway handler
	if err := pb.RegisterAuthHandlerServer(context.Background(), gwMux, s); err != nil {
		s.logger.Error("failed to register HTTP gateway", "err", err)
		return err
	}
	mux.Handle("/", gwMux)

	// --- Swagger UI ---
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./static/swagger-ui"))))

	// --- Swagger JSON ---
	mux.Handle("/swagger/", http.StripPrefix("/swagger", http.FileServer(http.Dir("docs/swagger"))))

	// HTTP сервер
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.cfg.HTTP.Ip, s.cfg.HTTP.Port),
		Handler: mux,
	}

	// Запуск HTTP сервера
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("failed to start HTTP gateway", "err", err)
		}
	}()
	s.httpListener = httpServer

	// --- gRPC Listener ---
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GRPC.Port))
	if err != nil {
		s.logger.Error("failed to listen", "err", err)
		return err
	}
	s.grpcListener = lis

	// Запуск gRPC сервера
	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			s.logger.Error("failed to start gRPC server", "err", err)
		}
	}()

	s.logger.Info("gRPC and HTTP servers started")
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.grpcServer.GracefulStop()

	grpcErr := s.grpcListener.Close()
	var httpErr error
	if s.httpListener != nil {
		httpErr = s.httpListener.Shutdown(ctx)
	}

	if httpErr != nil && grpcErr != nil {
		s.logger.Error("failed to stop server", "grpcErr", grpcErr, "httpErr", httpErr)
		return errors.New("failed to stop server")
	}

	s.logger.Info("gRPC and HTTP servers successfully stopped.")
	return nil
}
