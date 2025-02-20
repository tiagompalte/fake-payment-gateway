package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/tiagompalte/fake-payment-gateway/internal/app/entity"
	"github.com/tiagompalte/fake-payment-gateway/internal/app/protocols"
	"github.com/tiagompalte/fake-payment-gateway/pkg/cache"
	"github.com/tiagompalte/fake-payment-gateway/pkg/errors"
)

type FindAccountByTokenUseCase interface {
	Execute(ctx context.Context, token string) (entity.Account, error)
}

type FindAccountByTokenUseCaseImpl struct {
	account protocols.AccountRepository
	cache   cache.Cache
}

func NewFindAccountByTokenUseCaseImpl(account protocols.AccountRepository, cache cache.Cache) FindAccountByTokenUseCase {
	return FindAccountByTokenUseCaseImpl{
		account: account,
		cache:   cache,
	}
}

func (u FindAccountByTokenUseCaseImpl) Execute(ctx context.Context, token string) (entity.Account, error) {
	cacheKey := fmt.Sprintf("account_token:%s", token)

	var account entity.Account
	err := u.cache.Get(ctx, cacheKey, &account)
	if err == nil {
		return account, errors.Wrap(err)
	}

	account, err = u.account.FindByToken(ctx, token)
	if err != nil {
		return entity.Account{}, errors.Wrap(err)
	}

	u.cache.Set(ctx, cacheKey, account, time.Hour)

	return account, nil
}
