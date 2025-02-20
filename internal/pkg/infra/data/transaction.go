package data

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/tiagompalte/fake-payment-gateway/internal/app/entity"
	"github.com/tiagompalte/fake-payment-gateway/internal/app/protocols"
	"github.com/tiagompalte/fake-payment-gateway/pkg/errors"
	"github.com/tiagompalte/fake-payment-gateway/pkg/repository"
)

type TransactionRepository struct {
	conn         repository.ConnectorSql
	mainTable    string
	selectFields string
}

func NewTransactionRepository(conn repository.ConnectorSql) protocols.TransactionRepository {
	return TransactionRepository{
		conn:      conn,
		mainTable: "tb_transaction",
		selectFields: `
			id
			, created_at
			, updated_at
			, uuid
			, account_id
			, status
			, name
			, credit_card_number
			, credit_card_security_code
			, credit_card_expires
			, amount
			
			FROM tb_transaction
		`,
	}
}

func (r TransactionRepository) parseEntity(s repository.Scanner) (entity.Transaction, error) {
	var transaction entity.Transaction
	err := s.Scan(
		&transaction.ID,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
		&transaction.UUID,
		&transaction.AccountID,
		&transaction.Status,
		&transaction.Name,
		&transaction.CreditCardNumber,
		&transaction.CreditCardSecurityCode,
		&transaction.CreditCardExpires,
		&transaction.Amount,
	)
	if err != nil {
		return entity.Transaction{}, errors.Repo(err, r.mainTable)
	}

	return transaction, nil
}

func (r TransactionRepository) Insert(ctx context.Context, transaction entity.Transaction) (uint32, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return 0, errors.Repo(err, r.mainTable)
	}

	res, err := r.conn.Exec(ctx,
		`INSERT INTO tb_transaction (
			uuid 
			, account_id
			, status
			, name
			, credit_card_number
			, credit_card_security_code
			, credit_card_expires
			, amount
		) VALUES (?,?,?,?,?,?,?,?)`,
		uuid,
		transaction.AccountID,
		transaction.Status,
		transaction.Name,
		transaction.CreditCardNumber,
		transaction.CreditCardSecurityCode,
		transaction.CreditCardExpires,
		transaction.Amount,
	)
	if err != nil {
		return 0, errors.Repo(err, r.mainTable)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Repo(err, r.mainTable)
	}

	return uint32(id), nil
}

func (r TransactionRepository) FindByID(ctx context.Context, id uint32) (entity.Transaction, error) {
	query := `
		SELECT %s
			WHERE NOT deleted_at AND id = ?
	`

	q := fmt.Sprintf(query, r.selectFields)

	transaction, err := r.parseEntity(
		r.conn.QueryRow(
			ctx,
			q,
			id,
		))
	if err != nil {
		return entity.Transaction{}, errors.Repo(err, r.mainTable)
	}

	return transaction, nil
}

func (r TransactionRepository) FindByUUID(ctx context.Context, uuid string) (entity.Transaction, error) {
	query := `
		SELECT %s
			WHERE NOT deleted_at AND uuid = ?
	`

	q := fmt.Sprintf(query, r.selectFields)

	transaction, err := r.parseEntity(
		r.conn.QueryRow(
			ctx,
			q,
			uuid,
		))
	if err != nil {
		return entity.Transaction{}, errors.Repo(err, r.mainTable)
	}

	return transaction, nil
}

func (r TransactionRepository) FindByAccountID(ctx context.Context, accountID uint32) ([]entity.Transaction, error) {
	query := `
		SELECT %s
			WHERE NOT deleted_at AND account_id = ?
	`

	q := fmt.Sprintf(query, r.selectFields)

	result, err := r.conn.Query(
		ctx,
		q,
		accountID,
	)
	list, err := repository.ParseEntities(
		r.parseEntity,
		result,
		err,
	)
	if err != nil {
		return []entity.Transaction{}, errors.Repo(err, r.mainTable)
	}

	return list, nil
}
