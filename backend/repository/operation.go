package repository

import (
	"calculadora/backend/domain"
	"database/sql"
	"log"
	"time"
)

type OperationQuery interface {
	CreateOperation (operation domain.Operation, db *sql.DB) (domain.Operation, error)
	GetOperations (db *sql.DB) ([]domain.Operation, error) 
}

type operationQuery struct {}

func (o *operationQuery) CreateOperation (operation domain.Operation, db *sql.DB) (domain.Operation, error){
	var operationID int
	var result float64

	query := `INSERT INTO "operation"("numero1", "numero2", "opeador", "resultado")
	VALUES ($1, $2, $3, $4) RETURNING id, resultado`

	err := db.QueryRow(query, operation.Number1, operation.Number2, operation.Operator, operation.Result).Scan(&operationID, &result)

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

func (o *operationQuery) GetOperations (db *sql.DB) ([]domain.Operation, error) {
	query := `SELECT * FROM "operation" ORDER BY "dthora_calculo" DESC`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var operations []domain.Operation
	var operation domain.Operation
	var createdAt *time.Time

	for rows.Next() {
		err := rows.Scan(
			&operation.ID,
			&operation.Number1,
			&operation.Number2,
			&operation.Result,
			&createdAt,
			&operation.Operator,
		)

		operation.StringCreatedAt = createdAt.Format(time.RFC850)
		
		if err != nil {
			log.Fatal(err)
		}

		operations = append(operations, operation)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return operations, nil
}


