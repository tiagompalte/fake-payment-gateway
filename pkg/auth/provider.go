package auth

import "github.com/tiagompalte/fake-payment-gateway/configs"

func ProviderSet(
	config configs.Config,
) Auth {
	return NewJwtAuth(config)
}
