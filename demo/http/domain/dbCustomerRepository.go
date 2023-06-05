package domain

import (
	"database/sql"
	"learning-http/errs"
	"learning-http/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}

type CustomerRepositoryDb struct {
	db *sqlx.DB
}

func (cr *CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	selectSQL := "SELECT customer_id, name, date_of_birth, city, zipcode, status from customers"

	if err := cr.db.Select(&customers, selectSQL); err != nil {
		logger.Error("Error while querying customer table: " + err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return customers, nil
}

func (cr *CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	c := Customer{}
	selectSQL := "SELECT customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id = ?"
	err := cr.db.Get(&c, selectSQL, id)

	// https://stackoverflow.com/a/60123886
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Info("No rows found")
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while querying customer table: " + err.Error())
			return nil, errs.NewUnexpectedError(err.Error())
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbConn *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{db: dbConn}
}
