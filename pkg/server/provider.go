package server

import (
	"github.com/tiagompalte/fake-payment-gateway/configs"
)

func ProviderSet(
	config configs.Config,
) Server {
	return NewGoChiServer(config)
}
