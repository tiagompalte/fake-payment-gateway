package repository

import (
	"github.com/tiagompalte/fake-payment-gateway/configs"
)

func ProviderDataSqlManagerSet(
	config configs.Config,
) DataSqlManager {
	return NewDataSqlWithConfig(config.Database)
}

func ProviderConnectorSqlSet(
	config configs.Config,
) ConnectorSql {
	dataSql := ProviderDataSqlManagerSet(config)
	return dataSql.Command()
}
