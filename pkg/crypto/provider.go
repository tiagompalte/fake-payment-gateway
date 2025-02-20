package crypto

import "github.com/tiagompalte/fake-payment-gateway/configs"

func ProviderSet(
	config configs.Config,
) Crypto {
	return NewBcrypt(config.Bcrypt)
}
