package service

import (
	"database/sql"
	"log"
	"calculadora/backend/domain"
	"calculadora/backend/repository"
)

type OperationService interface {
	CreateOperation (operation domain.Operation, db *sql.DB) (domain.Operation, error)
	GetOperations (db *sql.DB) ([]domain.Operation, error) 
}

type operationService struct {
	dao repository.DAO
}

func NewOperationService(dao repository.DAO) OperationService {
	return &operationService{dao: dao}
}

func (o *operationService) CreateOperation (operation domain.Operation, db *sql.DB) (domain.Operation, error) {
	operationInfo, err := o.dao.NewOperationQuery().CreateOperation(operation, db)
	
	if err != nil {
		log.Fatal(err)
		return operationInfo, err
	}

	return operationInfo, nil
}

func (o *operationService) GetOperations (db *sql.DB) ([]domain.Operation, error) {
	operations, err := o.dao.NewOperationQuery().GetOperations(db)

	if err != nil {
		log.Fatal(err)
	}
	
	return operations, nil
}
