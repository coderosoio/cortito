package wrapper

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"

	"common/auth"
)

const (
	methodKey = "Micro-Method"
)

var (
	// TODO: This should be configurable
	allowed = []string{
		"User.CreateUser",
		"Auth.CreateToken",
		// "Auth.VerifyToken",
	}
)

func NewAuthCallWrapper() client.CallWrapper {
	return func(cf client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node *registry.Node, req client.Request, res interface{}, opts client.CallOptions) error {
			md, ok := metadata.FromContext(ctx)
			if !ok {
				md = map[string]string{}
			}
			if token := ctx.Value("token"); token != nil {
				md["authorization"] = token.(string)
			}
			ctx = metadata.NewContext(ctx, md)
			return cf(ctx, node, req, res, opts)
		}
	}
}

func NewAuthHandlerWrapper(authStrategy auth.Auth) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, res interface{}) error {
			md, ok := metadata.FromContext(ctx)
			if !ok {
				md = make(map[string]string)
			}
			md = metadata.Copy(md)
			method := md[methodKey]

			if !isAllowed(method) {
				token, found := md["Authorization"]
				if !found {
					return fmt.Errorf("not authorized call")
				}
				if _, err := authStrategy.VerifyToken(token); err != nil {
					return fmt.Errorf("invalid token: %v", err)
				}
			}
			return h(ctx, req, res)
		}
	}
}

func isAllowed(method string) bool {
	for _, allowed := range allowed {
		if method == allowed {
			return true
		}
	}
	return false
}
