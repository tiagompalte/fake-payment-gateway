package job

import (
	"context"
	"log"

	"github.com/tiagompalte/fake-payment-gateway/internal/app/protocols"
	"github.com/tiagompalte/fake-payment-gateway/internal/app/usecase"
)

const CreateAccountJobName = "create-account"

type createAccountJobImpl struct {
	createAccountUseCase usecase.CreateAccountUseCase
}

func NewCreateAccountJobImpl(createAccountUseCase usecase.CreateAccountUseCase) protocols.Job {
	return createAccountJobImpl{
		createAccountUseCase,
	}
}

func (j createAccountJobImpl) Execute(ctx context.Context, args ...any) error {
	account, err := j.createAccountUseCase.Execute(context.Background())
	if err != nil {
		log.Fatalf("failed to create account (error: %v)", err)
	}

	log.Printf("%+v", account)

	return nil
}
