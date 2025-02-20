package main

import (
	"context"
	"log"

	"github.com/tiagompalte/fake-payment-gateway/application"
)

func main() {
	app, err := application.Build()
	if err != nil {
		log.Fatalf("failed to build the application (error: %v)", err)
	}

	account, err := app.UseCase().CreateAccountUseCase.Execute(context.Background())
	if err != nil {
		log.Fatalf("failed to create account (error: %v)", err)
	}

	log.Printf("%+v", account)
}
