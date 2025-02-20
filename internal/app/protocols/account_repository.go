package protocols

import (
	"context"

	"github.com/tiagompalte/fake-payment-gateway/internal/app/entity"
)

type AccountRepository interface {
	Insert(ctx context.Context, account entity.Account) (uint32, error)
	FindByID(ctx context.Context, id uint32) (entity.Account, error)
	FindByToken(ctx context.Context, token string) (entity.Account, error)
}
