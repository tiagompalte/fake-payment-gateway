package protocols

import (
	"context"

	"github.com/tiagompalte/fake-payment-gateway/internal/app/entity"
)

type TransactionRepository interface {
	Insert(ctx context.Context, transaction entity.Transaction) (uint32, error)
	FindByID(ctx context.Context, id uint32) (entity.Transaction, error)
	FindByUUID(ctx context.Context, uuid string) (entity.Transaction, error)
	FindByAccountID(ctx context.Context, accountID uint32) ([]entity.Transaction, error)
}
