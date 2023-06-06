package domain

import (
	"database/sql"
	"learning-http-auth/errs"
	"learning-http-auth/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type AuthRepositoryDb struct {
	db *sqlx.DB
}

func (d AuthRepositoryDb) FindBy(username, password string) *errs.AppError {
	var login string
	sqlVerify := `SELECT username FROM users u WHERE username = ? and password = ?`
	err := d.db.Get(&login, sqlVerify, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewAuthenticationError("invalid credentials")
		} else {
			logger.Error("Error while verifying login request from database: " + err.Error())
			return errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return nil
}

func NewAuthRepository(client *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client}
}
