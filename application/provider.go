package application

import (
	"github.com/google/wire"
	"github.com/tiagompalte/fake-payment-gateway/internal/app/usecase"
	"github.com/tiagompalte/fake-payment-gateway/internal/pkg/infra"
	"github.com/tiagompalte/fake-payment-gateway/pkg"
)

var ProviderSet = wire.NewSet(
	pkg.ProviderSet,
	infra.ProviderSet,
	usecase.ProviderSet,
	wire.Struct(new(usecase.UseCase), "*"),
	ProvideApplication,
)
