package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/tiagompalte/fake-payment-gateway/internal/app/entity"
	"github.com/tiagompalte/fake-payment-gateway/internal/app/protocols"
	"github.com/tiagompalte/fake-payment-gateway/pkg/errors"
)

type CreateAccountUseCase interface {
	Execute(ctx context.Context) (entity.Account, error)
}

type createAccountUseCaseImpl struct {
	accountRepository protocols.AccountRepository
}

func NewCreateAccountUseCaseImpl(accountRepository protocols.AccountRepository) CreateAccountUseCase {
	return createAccountUseCaseImpl{
		accountRepository,
	}
}

func (u createAccountUseCaseImpl) Execute(ctx context.Context) (entity.Account, error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return entity.Account{}, errors.Wrap(err)
	}

	var account entity.Account
	account.Token = token.String()

	id, err := u.accountRepository.Insert(ctx, account)
	if err != nil {
		return entity.Account{}, errors.Wrap(err)
	}

	account, err = u.accountRepository.FindByID(ctx, id)
	if err != nil {
		return entity.Account{}, errors.Wrap(err)
	}

	return account, nil
}
