package routes

import (
	"github.com/tiagompalte/fake-payment-gateway/application"
	"github.com/tiagompalte/fake-payment-gateway/pkg/server"
)

func CreateRouteV1(app application.App) server.GroupRoute {
	return server.GroupRoute{
		Path: "/v1",
		GroupRoutes: []server.GroupRoute{
			CreateGroupTransactionV1(app),
		},
	}
}
