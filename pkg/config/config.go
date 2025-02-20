package config

import "github.com/tiagompalte/fake-payment-gateway/configs"

type Config interface {
	Load(filename string, configType string, path string) (configs.Config, error)
}
