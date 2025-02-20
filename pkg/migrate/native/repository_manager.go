package nativemigrate

import "github.com/tiagompalte/fake-payment-gateway/pkg/repository"

type repo struct {
	data repository.DataSqlManager
}

func NewRepositoryManager(data repository.DataSqlManager) RepositoryManager {
	return repo{data: data}
}

func (r repo) Data() repository.DataSqlManager {
	return r.data
}

func (r repo) Version(conn repository.ConnectorSql) VersionRepository {
	return NewVersionRepositoryImpl(conn)
}
