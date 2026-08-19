package users_transport_http

import (
	"net/http"

	core_logger "github.com/Crysta1l/go-tdApp/internal/core/logger"
	core_http_response "github.com/Crysta1l/go-tdApp/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
}
