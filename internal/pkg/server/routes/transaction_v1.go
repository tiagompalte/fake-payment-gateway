package routes

import (
	"github.com/tiagompalte/fake-payment-gateway/application"
	"github.com/tiagompalte/fake-payment-gateway/internal/pkg/server/constant"
	"github.com/tiagompalte/fake-payment-gateway/internal/pkg/server/handler"
	"github.com/tiagompalte/fake-payment-gateway/internal/pkg/server/middleware"
	"github.com/tiagompalte/fake-payment-gateway/pkg/server"
)

func CreateGroupTransactionV1(app application.App) server.GroupRoute {
	routes := []server.Route{
		{
			Path:    "/",
			Method:  server.RouteMethodPost,
			Handler: handler.CreateTransactionHandler(app.UseCase().CreateTransactionUseCase()),
		},
	}

	return server.GroupRoute{
		Path: "/transactions",
		Middlewares: []server.Middleware{
			middleware.ValidateExtractAccountTokenMiddleware(constant.AccessToken, app.UseCase().FindAccountByTokenUseCase()),
		},
		Routes: routes,
	}
}
