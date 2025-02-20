package usecase

import (
	"context"
	"time"

	"github.com/tiagompalte/fake-payment-gateway/internal/app/entity"
	"github.com/tiagompalte/fake-payment-gateway/internal/app/protocols"
	"github.com/tiagompalte/fake-payment-gateway/pkg/errors"
)

type CreateTransactionUseCase interface {
	Execute(ctx context.Context, input CreateTransactionInput) (entity.Transaction, error)
}

type CreateTransactionInput struct {
	AccountID              uint32
	Name                   string
	CreditCardNumber       string
	CreditCardSecurityCode string
	CreditCardExpiresYear  int
	CreditCardExpiresMonth int
	Amount                 float64
}

type CreateTransactionUseCaseImpl struct {
	transactionRepository protocols.TransactionRepository
}

func NewCreateTransactionUseCaseImpl(transactionRepository protocols.TransactionRepository) CreateTransactionUseCase {
	return CreateTransactionUseCaseImpl{
		transactionRepository: transactionRepository,
	}
}

func (u CreateTransactionUseCaseImpl) Execute(ctx context.Context, input CreateTransactionInput) (entity.Transaction, error) {
	var transaction entity.Transaction
	transaction.AccountID = input.AccountID
	transaction.Status = entity.TransactionStatusApproved
	transaction.Name = input.Name
	transaction.CreditCardNumber = input.CreditCardNumber
	transaction.CreditCardSecurityCode = input.CreditCardSecurityCode
	transaction.CreditCardExpires = time.Date(input.CreditCardExpiresYear, time.Month(input.CreditCardExpiresMonth), 1, 0, 0, 0, 0, time.UTC)
	transaction.Amount = input.Amount

	id, err := u.transactionRepository.Insert(ctx, transaction)
	if err != nil {
		return entity.Transaction{}, errors.Wrap(err)
	}

	transaction, err = u.transactionRepository.FindByID(ctx, id)
	if err != nil {
		return entity.Transaction{}, errors.Wrap(err)
	}

	return transaction, nil
}
