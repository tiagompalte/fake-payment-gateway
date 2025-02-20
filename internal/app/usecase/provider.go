package usecase

import (
	"github.com/google/wire"
	"github.com/tiagompalte/fake-payment-gateway/pkg/cache"
	"github.com/tiagompalte/fake-payment-gateway/pkg/healthcheck"
	"github.com/tiagompalte/fake-payment-gateway/pkg/repository"
)

var ProviderSet = wire.NewSet(
	ProviderHealthCheckUseCase,
	NewFindAccountByTokenUseCaseImpl,
	NewCreateTransactionUseCaseImpl,
	NewCreateAccountUseCaseImpl,
)

func ProviderHealthCheckUseCase(cache cache.Cache, dataSqlManager repository.DataSqlManager) HealthCheckUseCase {
	return NewHealthCheckUseCaseImpl([]healthcheck.HealthCheck{
		cache, dataSqlManager,
	})
}
