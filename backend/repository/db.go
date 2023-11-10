package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = "5432"
	dbname = "db_calculator"
	user = "postgres"
	password = "romeulindo"
)

var DB *sql.DB

type DAO interface{
	NewOperationQuery() OperationQuery
}

type dao struct {}

func NewDAO() DAO {
	return &dao{}
}

func NewDB() (*sql.DB, error) {
	dbconn := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable user=%s password=%s", host, port, dbname, user, password)
	DB, err := sql.Open("postgres", dbconn)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return DB, nil
}

func (d *dao) NewOperationQuery() OperationQuery {
	return &operationQuery{}
}
