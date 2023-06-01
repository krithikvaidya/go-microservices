package domain

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	db *sql.DB
}

func (cr *CustomerRepositoryDb) FindAll() ([]Customer, error) {
	selectSQL := "SELECT customer_id, name, date_of_birth, city, zipcode, status from customers"
	rows, err := cr.db.Query(selectSQL)
	if err != nil {
		log.Println("Error while querying customer table: ", err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)

	for rows.Next() {
		c := Customer{}
		err = rows.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error while scanning customer data: ", err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (cr *CustomerRepositoryDb) ById(id string) (*Customer, error) {
	selectSQL := "SELECT customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id = ?"
	row := cr.db.QueryRow(selectSQL, id)

	if row.Err() != nil {
		log.Println("Error while querying customer table: ", row.Err())
		return nil, row.Err()
	}

	c := Customer{}
	err := row.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)

	// https://stackoverflow.com/a/60123886
	if err == sql.ErrNoRows {
		log.Println("No rows found...", row.Err())
		return nil, errors.New("Customer not found")
	}
	if err != nil {
		log.Println("Error while scanning customer data: ", err.Error())
		return nil, err
	}

	return &c, nil
}

func getDbClient() *sql.DB {
	db, err := sql.Open("mysql", "root:student@tcp(localhost:3307)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	return CustomerRepositoryDb{db: getDbClient()}
}