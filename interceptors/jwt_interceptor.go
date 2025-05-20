package interceptors

import (
	"context"
	"final/internal/security/jwt"
	"final/internal/utils/observability/log"
	"fmt"
	"google.golang.org/grpc"
	grpc_health_v1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
)

type claimsKey string

const ClaimsKey claimsKey = claimsKey("claims")

func JwtInterceptor(logger *log.Logger, enabled bool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !enabled || info.FullMethod == grpc_health_v1.Health_Check_FullMethodName {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			logger.Error("missing metadata")
			return nil, fmt.Errorf("missing metadata")
		}

		authHeader := md["authorization"]
		if len(authHeader) == 0 {
			logger.Error("missing authorization token")
			return nil, fmt.Errorf("missing authorization token")
		}

		accessToken := authHeader[0]

		claims, err := jwt.DecodeAccessToken(accessToken)
		if err != nil {
			logger.Error("failed to decode token", "err", err)
			return nil, fmt.Errorf("failed to decode token: %s", err)
		}

		ctx = context.WithValue(ctx, ClaimsKey, claims)
		return handler(ctx, req)
	}
}
