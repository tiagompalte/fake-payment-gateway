package infra

import (
	"github.com/google/wire"
	"github.com/tiagompalte/fake-payment-gateway/internal/pkg/infra/data"
	"github.com/tiagompalte/fake-payment-gateway/internal/pkg/infra/uow"
)

var ProviderSet = wire.NewSet(
	data.ProviderSet,
	uow.ProviderSet,
)
