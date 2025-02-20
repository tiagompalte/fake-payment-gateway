package pkg

import (
	"github.com/google/wire"
	"github.com/tiagompalte/fake-payment-gateway/pkg/auth"
	"github.com/tiagompalte/fake-payment-gateway/pkg/cache"
	"github.com/tiagompalte/fake-payment-gateway/pkg/config"
	"github.com/tiagompalte/fake-payment-gateway/pkg/crypto"
	"github.com/tiagompalte/fake-payment-gateway/pkg/log"
	"github.com/tiagompalte/fake-payment-gateway/pkg/repository"
	"github.com/tiagompalte/fake-payment-gateway/pkg/server"
)

var ProviderSet = wire.NewSet(
	config.ProviderSet,
	cache.ProviderSet,
	repository.ProviderDataSqlManagerSet,
	repository.ProviderConnectorSqlSet,
	server.ProviderSet,
	crypto.ProviderSet,
	auth.ProviderSet,
	log.ProviderSet,
)
