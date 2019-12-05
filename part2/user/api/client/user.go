package client

import (
	"context"
	"github.com/entere/micro-examples/part2/user/srv/proto/user"
	"github.com/micro/go-micro"

	"github.com/micro/go-micro/server"
)

type userKey struct{}

// FromContext retrieves the client from the Context
func UserFromContext(ctx context.Context) (io_github_entere_srv_user.UserService, bool) {
	c, ok := ctx.Value(userKey{}).(io_github_entere_srv_user.UserService)
	return c, ok
}

// Client returns a wrapper for the UserClient
func UserWrapper(service micro.Service) server.HandlerWrapper {
	client := io_github_entere_srv_user.NewUserService("io.github.entere.srv.student", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, userKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
