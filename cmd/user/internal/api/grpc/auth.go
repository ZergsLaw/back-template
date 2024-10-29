package grpc

import (
	"context"
	"fmt"
	"github.com/ZergsLaw/back-template/cmd/user/internal/session"
	"path"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"

	"github.com/ZergsLaw/back-template/cmd/user/internal/app"
)

// AuthFuncOverride implements grpc_auth.ServiceAuthFuncOverride.
func (a *api) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	_, method := path.Split(fullMethodName)
	if !a.auth[method] {
		return ctx, nil
	}

	token, err := grpc_auth.AuthFromMD(ctx, scheme)
	if err != nil {
		return nil, fmt.Errorf("grpc_auth.AuthFromMD: %w", err)
	}

	userSession, err := a.app.Auth(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", app.ErrInvalidAuth, err)
	}

	return session.NewContext(ctx, userSession), nil
}
