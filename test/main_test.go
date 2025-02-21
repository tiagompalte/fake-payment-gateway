//go:build integration
// +build integration

package test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/tiagompalte/fake-payment-gateway/internal/app/entity"
	"github.com/tiagompalte/fake-payment-gateway/test/testconfig"
)

func TestMain(t *testing.M) {
	config := testconfig.Instance()

	defer config.Close()

	code := t.Run()

	os.Exit(code)
}

func GenerateAccountAndToken() entity.Account {
	ctx := context.Background()

	app := testconfig.Instance().App()

	account, err := app.UseCase().CreateAccountUseCase().Execute(ctx)
	if err != nil {
		log.Fatalf("failed to create account: %v", err)
	}

	return account
}
