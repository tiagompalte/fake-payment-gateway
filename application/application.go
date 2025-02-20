package application

import (
	"github.com/tiagompalte/fake-payment-gateway/configs"
	"github.com/tiagompalte/fake-payment-gateway/internal/app/usecase"
	"github.com/tiagompalte/fake-payment-gateway/pkg/auth"
	"github.com/tiagompalte/fake-payment-gateway/pkg/log"
	"github.com/tiagompalte/fake-payment-gateway/pkg/server"
)

type App struct {
	config  configs.Config
	server  server.Server
	useCase usecase.UseCase
	auth    auth.Auth
	log     log.Log
}

func ProvideApplication(
	config configs.Config,
	server server.Server,
	useCase usecase.UseCase,
	auth auth.Auth,
	log log.Log,
) App {
	return App{
		config,
		server,
		useCase,
		auth,
		log,
	}
}

func (app App) Config() configs.Config {
	return app.config
}

func (app App) Server() server.Server {
	return app.server
}

func (app App) UseCase() usecase.UseCase {
	return app.useCase
}

func (app App) Auth() auth.Auth {
	return app.auth
}

func (app App) Log() log.Log {
	return app.log
}
