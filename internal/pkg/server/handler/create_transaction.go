package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tiagompalte/fake-payment-gateway/internal/app/entity"
	"github.com/tiagompalte/fake-payment-gateway/internal/app/usecase"
	"github.com/tiagompalte/fake-payment-gateway/internal/pkg/server/constant"
	"github.com/tiagompalte/fake-payment-gateway/pkg/errors"
	"github.com/tiagompalte/fake-payment-gateway/pkg/server"
)

type CreateTransactionRequest struct {
	Name                   string  `json:"name"`
	CreditCardNumber       string  `json:"credit_card_number"`
	CreditCardSecurityCode string  `json:"credit_card_security_code"`
	CreditCardExpiresYear  int     `json:"credit_card_expires_year"`
	CreditCardExpiresMonth int     `json:"credit_card_expires_month"`
	Amount                 float64 `json:"amount"`
}

func (r CreateTransactionRequest) toInput() usecase.CreateTransactionInput {
	return usecase.CreateTransactionInput{
		Name:                   r.Name,
		CreditCardNumber:       r.CreditCardNumber,
		CreditCardSecurityCode: r.CreditCardSecurityCode,
		CreditCardExpiresYear:  r.CreditCardExpiresYear,
		CreditCardExpiresMonth: r.CreditCardExpiresMonth,
		Amount:                 r.Amount,
	}
}

type CreateTransactionResponse struct {
	UUID   string `json:"uuid"`
	Status string `json:"status"`
}

// @Summary Create Transaction
// @Description Create new transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Security apiKey
// @Param new_task body CreateTransactionRequest true "New Transaction"
// @Success 201 {object} CreateTransactionResponse "Create Transaction success"
// @Router /api/v1/transactions [post]
func CreateTransactionHandler(createTransactionUseCase usecase.CreateTransactionUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var request CreateTransactionRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return errors.Wrap(err)
		}

		account, ok := ctx.Value(constant.ContextAccount).(entity.Account)
		if !ok {
			return errors.Wrap(errors.NewInvalidTokenError())
		}

		input := request.toInput()
		input.AccountID = account.ID

		transaction, err := createTransactionUseCase.Execute(ctx, input)
		if err != nil {
			return errors.Wrap(err)
		}

		resp := CreateTransactionResponse{
			UUID:   transaction.UUID,
			Status: transaction.Status.String(),
		}

		err = server.RespondJSON(w, http.StatusCreated, resp)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
