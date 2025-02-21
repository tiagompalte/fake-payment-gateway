// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package application

import (
	"github.com/tiagompalte/fake-payment-gateway/internal/app/usecase"
	"github.com/tiagompalte/fake-payment-gateway/internal/pkg/infra/data"
	"github.com/tiagompalte/fake-payment-gateway/pkg/auth"
	"github.com/tiagompalte/fake-payment-gateway/pkg/cache"
	"github.com/tiagompalte/fake-payment-gateway/pkg/config"
	"github.com/tiagompalte/fake-payment-gateway/pkg/log"
	"github.com/tiagompalte/fake-payment-gateway/pkg/repository"
	"github.com/tiagompalte/fake-payment-gateway/pkg/server"
)

// Injectors from wire.go:

func Build() (App, error) {
	configsConfig := config.ProviderSet()
	serverServer := server.ProviderSet(configsConfig)
	cacheCache := cache.ProviderSet(configsConfig)
	dataManager := repository.ProviderDataSqlManagerSet(configsConfig)
	healthCheckUseCase := usecase.ProviderHealthCheckUseCase(cacheCache, dataManager)
	connectorSql := repository.ProviderConnectorSqlSet(configsConfig)
	accountRepository := data.NewAccountRepository(connectorSql)
	findAccountByTokenUseCase := usecase.NewFindAccountByTokenUseCaseImpl(accountRepository, cacheCache)
	transactionRepository := data.NewTransactionRepository(connectorSql)
	createTransactionUseCase := usecase.NewCreateTransactionUseCaseImpl(transactionRepository, configsConfig)
	createAccountUseCase := usecase.NewCreateAccountUseCaseImpl(accountRepository)
	useCase := usecase.NewUseCase(healthCheckUseCase, findAccountByTokenUseCase, createTransactionUseCase, createAccountUseCase)
	authAuth := auth.ProviderSet(configsConfig)
	logLog := log.ProviderSet()
	app := ProvideApplication(configsConfig, serverServer, useCase, authAuth, logLog)
	return app, nil
}
