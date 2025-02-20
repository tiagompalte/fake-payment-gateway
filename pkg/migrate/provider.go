package migrate

import (
	"github.com/tiagompalte/fake-payment-gateway/configs"
	nativemigrate "github.com/tiagompalte/fake-payment-gateway/pkg/migrate/native"
	"github.com/tiagompalte/fake-payment-gateway/pkg/repository"
)

func ProviderSet(
	config configs.Config,
	data repository.DataSqlManager,
) Migrate {
	if config.Migrate.DriverName == configs.GolangMigrate {
		return NewGolangMigrate(config)
	} else if config.Migrate.DriverName == configs.NativeMigrate {
		file := nativemigrate.NewFileImpl(config.Migrate.PathMigrations)
		repositoryManager := nativemigrate.NewRepositoryManager(data)
		return nativemigrate.NewNativeMigrate(file, repositoryManager)
	}
	panic("None migrate define")
}
