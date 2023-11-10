package repository

import (
	"database/sql"
	"log"
	"calculadora/domain"
)

type OperationQuery interface {
	CreateOperation (operation domain.Operation, db *sql.DB) (domain.Operation, error)
}

type operationQuery struct {}

func (o *operationQuery) CreateOperation (operation domain.Operation, db *sql.DB) (domain.Operation, error){
	var operationID int
	var result float64

	query := `INSERT INTO "operation"("numero1", "numero2", "operador", "resultado")
	VALUES ($1, $2, $3) RETURNING id, resultado`

	err := db.QueryRow(query, operation.ID, operation.Number1, operation.Number2, operation.Operator, operation.Result).Scan(&operationID, &result)

	operationInfo := domain.Operation{
		ID: operationID,
		Result: result,
	}
	
	if err != nil {
		log.Fatal(err)
		return domain.Operation{}, err
	}

	return operationInfo, nil
}


