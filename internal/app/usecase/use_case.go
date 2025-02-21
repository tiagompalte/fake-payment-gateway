package usecase

type UseCase struct {
	healthCheckUseCase        HealthCheckUseCase
	findAccountByTokenUseCase FindAccountByTokenUseCase
	createTransactionUseCase  CreateTransactionUseCase
	createAccountUseCase      CreateAccountUseCase
}

func NewUseCase(
	healthCheckUseCase HealthCheckUseCase,
	findAccountByTokenUseCase FindAccountByTokenUseCase,
	createTransactionUseCase CreateTransactionUseCase,
	createAccountUseCase CreateAccountUseCase,
) UseCase {
	return UseCase{
		healthCheckUseCase,
		findAccountByTokenUseCase,
		createTransactionUseCase,
		createAccountUseCase,
	}
}

func (u UseCase) HealthCheckUseCase() HealthCheckUseCase {
	return u.healthCheckUseCase
}

func (u UseCase) FindAccountByTokenUseCase() FindAccountByTokenUseCase {
	return u.findAccountByTokenUseCase
}

func (u UseCase) CreateTransactionUseCase() CreateTransactionUseCase {
	return u.createTransactionUseCase
}

func (u UseCase) CreateAccountUseCase() CreateAccountUseCase {
	return u.createAccountUseCase
}
