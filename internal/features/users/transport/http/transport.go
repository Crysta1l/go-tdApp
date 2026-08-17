package users_transport_http

import (
	"context"
	"net/http"

	"github.com/Crysta1l/go-tdApp/internal/core/domain"
	core_http_server "github.com/Crysta1l/go-tdApp/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

// Fixed err with handleFunc
func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: http.HandlerFunc(h.CreateUser),
		},
	}
}
