package domain

import (
	"database/sql"
	"time"
	"github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	client, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	findAllSql := "select customer_id, city, zipcode, date of birth, status from customers"

	rows, err := client.Query(findAllSql)

	if err != nil {
		log.Prinln("Error while querying customers table "+ err.Error())
		
	}
}
