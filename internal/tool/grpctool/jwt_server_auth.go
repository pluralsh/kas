package grpctool

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/tool/logz"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JWTAuther struct {
	jwtIssuer         string
	jwtAudience       string
	secret            []byte
	loggerFromContext func(context.Context) *zap.Logger
}

func NewJWTAuther(secret []byte, jwtIssuer, jwtAudience string, loggerFromContext func(context.Context) *zap.Logger) *JWTAuther {
	return &JWTAuther{
		jwtIssuer:         jwtIssuer,
		jwtAudience:       jwtAudience,
		secret:            secret,
		loggerFromContext: loggerFromContext,
	}
}

// UnaryServerInterceptor returns a new unary server interceptors that performs per-request JWT auth.
func (a *JWTAuther) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := a.doAuth(ctx); err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

// StreamServerInterceptor returns a new stream server interceptors that performs per-request JWT auth.
func (a *JWTAuther) StreamServerInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := a.doAuth(stream.Context()); err != nil {
		return err
	}
	return handler(srv, stream)
}

func (a *JWTAuther) doAuth(ctx context.Context) error {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return err // returns gRPC status error
	}
	var claims jwt.RegisteredClaims
	_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return a.secret, nil
	}, jwt.WithValidMethods([]string{"HS256", "HS384", "HS512"}))
	if err != nil {
		a.loggerFromContext(ctx).Debug("JWT auth failed", logz.Error(err))
		return status.Error(codes.Unauthenticated, "JWT validation failed")
	}
	if a.jwtAudience != "" {
		if !claims.VerifyAudience(a.jwtAudience, true) {
			return status.Error(codes.Unauthenticated, "JWT audience validation failed")
		}
	}
	if a.jwtIssuer != "" {
		if !claims.VerifyIssuer(a.jwtIssuer, true) {
			return status.Error(codes.Unauthenticated, "JWT issuer validation failed")
		}
	}

	return nil
}
