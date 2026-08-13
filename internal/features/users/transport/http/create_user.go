package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Crysta1l/go-tdApp/internal/core/logger"
	core_http_request "github.com/Crysta1l/go-tdApp/internal/core/transport/http/request"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required, min=3, max=100,"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty, min=10, max=15, startswith=+"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	log.Debug("Invoke CreateUser handler")

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		fmt.Errorf("Create user %w", err)
	}

	rw.WriteHeader(http.StatusOK)
}
