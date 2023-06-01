package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct{}

func (cr *CustomerRepositoryDb) FindAll() ([]Customer, error) {
	db := getDbClient()

	selectSQL := "SELECT customer_id, name, date_of_birth, city, zipcode, status from customers"
	rows, err := db.Query(selectSQL)
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
	return CustomerRepositoryDb{}
}
