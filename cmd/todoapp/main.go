package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/Crysta1l/go-tdApp/internal/core/logger"
	core_http_middleware "github.com/Crysta1l/go-tdApp/internal/core/transport/http/middleware"
	core_http_server "github.com/Crysta1l/go-tdApp/internal/core/transport/http/server"
	users_transport_http "github.com/Crysta1l/go-tdApp/internal/features/users/transport/http"
	"go.uber.org/zap"
)

// sudo chown -R $USER:$USER .
func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)

	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("Failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Starting todo application")

	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(nil)

	usersRoutes := usersTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRouter(usersRoutes...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}

}
