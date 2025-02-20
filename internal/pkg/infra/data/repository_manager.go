package data

import (
	"github.com/tiagompalte/fake-payment-gateway/internal/app/protocols"
	"github.com/tiagompalte/fake-payment-gateway/pkg/repository"
)

type RepositoryManager interface {
	Account() protocols.AccountRepository
	Transaction() protocols.TransactionRepository
}

type repo struct {
	account     protocols.AccountRepository
	transaction protocols.TransactionRepository
}

func NewRepositoryManager(conn repository.ConnectorSql) RepositoryManager {
	return repo{
		account:     NewAccountRepository(conn),
		transaction: NewTransactionRepository(conn),
	}
}

func (r repo) Account() protocols.AccountRepository {
	return r.account
}

func (r repo) Transaction() protocols.TransactionRepository {
	return r.transaction
}
