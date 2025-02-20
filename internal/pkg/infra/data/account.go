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

type AccountRepository struct {
	conn         repository.ConnectorSql
	mainTable    string
	selectFields string
}

func NewAccountRepository(conn repository.ConnectorSql) protocols.AccountRepository {
	return AccountRepository{
		conn:      conn,
		mainTable: "tb_account",
		selectFields: `
			id
			, created_at
			, updated_at
			, uuid
			, token
			
			FROM tb_account
		`,
	}
}

func (r AccountRepository) parseEntity(s repository.Scanner) (entity.Account, error) {
	var account entity.Account
	err := s.Scan(
		&account.ID,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.UUID,
		&account.Token,
	)
	if err != nil {
		return entity.Account{}, errors.Repo(err, r.mainTable)
	}

	return account, nil
}

func (r AccountRepository) Insert(ctx context.Context, account entity.Account) (uint32, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return 0, errors.Repo(err, r.mainTable)
	}

	res, err := r.conn.Exec(ctx,
		"INSERT INTO tb_account (uuid, token) VALUES (?,?)",
		uuid,
		account.Token,
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

func (r AccountRepository) FindByID(ctx context.Context, id uint32) (entity.Account, error) {
	query := `
		SELECT %s
			WHERE NOT deleted_at AND id = ?
	`

	q := fmt.Sprintf(query, r.selectFields)

	account, err := r.parseEntity(
		r.conn.QueryRow(
			ctx,
			q,
			id,
		))
	if err != nil {
		return entity.Account{}, errors.Repo(err, r.mainTable)
	}

	return account, nil
}

func (r AccountRepository) FindByToken(ctx context.Context, token string) (entity.Account, error) {
	query := `
		SELECT %s
			WHERE NOT deleted_at AND token = ?
	`

	q := fmt.Sprintf(query, r.selectFields)

	account, err := r.parseEntity(
		r.conn.QueryRow(
			ctx,
			q,
			token,
		))
	if err != nil {
		return entity.Account{}, errors.Repo(err, r.mainTable)
	}

	return account, nil
}
