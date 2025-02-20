package usecase

import (
	"context"
	"slices"
	"strings"
	"time"

	"github.com/tiagompalte/fake-payment-gateway/configs"
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
	transactionRepository   protocols.TransactionRepository
	creditCardNumbersDenied []string
}

func NewCreateTransactionUseCaseImpl(transactionRepository protocols.TransactionRepository, config configs.Config) CreateTransactionUseCase {
	creditCardNumbersDenied := make([]string, 0)
	if config.TransactionDenied.CreditCardNumbers != "" {
		creditCardNumbersDenied = strings.Split(config.TransactionDenied.CreditCardNumbers, ";")
	}

	return CreateTransactionUseCaseImpl{
		transactionRepository:   transactionRepository,
		creditCardNumbersDenied: creditCardNumbersDenied,
	}
}

func (u CreateTransactionUseCaseImpl) Execute(ctx context.Context, input CreateTransactionInput) (entity.Transaction, error) {
	status := entity.TransactionStatusApproved
	if slices.Contains(u.creditCardNumbersDenied, input.CreditCardNumber) {
		status = entity.TransactionStatusDenied
	}

	var transaction entity.Transaction
	transaction.AccountID = input.AccountID
	transaction.Status = status
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
