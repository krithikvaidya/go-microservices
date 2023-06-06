package domain

import (
	"learning-http-auth/errs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type AuthRepositoryDb struct {
	db *sqlx.DB
}

func (d AuthRepositoryDb) FindBy(username, password string) *errs.AppError {
	return nil
}

func NewAuthRepository(client *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client}
}
