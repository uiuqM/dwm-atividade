package service

import (
	"database/sql"
	"log"
	"calculadora/domain"
	"calculadora/repository"
)

type OperationService interface {
	CreateOperation (operation domain.Operation, db *sql.DB) (domain.Operation, error)
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
