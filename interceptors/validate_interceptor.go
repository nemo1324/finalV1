package interceptors

import (
	"context"
	"os"

	"final/internal/utils/observability/log"
	govalidator "github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func WithValidation(logger *log.Logger) grpc.UnaryServerInterceptor {
	v, err := govalidator.New()
	if err != nil {
		logger.Error("init govalidator failed", "err", err)
		os.Exit(1)
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		protoMsg, ok := req.(proto.Message)
		if !ok {
			logger.Warn("not valid proto message", "method", info.FullMethod)
			return nil, status.Errorf(codes.InvalidArgument, "not valid proto message")
		}

		if err = v.Validate(protoMsg); err != nil {
			logger.Warn("validation failed", "method", info.FullMethod, "err", err)
			return nil, status.Errorf(codes.InvalidArgument, "validation failed: %s", err)
		}

		return handler(ctx, req)
	}
}
