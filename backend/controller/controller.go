package controller

import (
	"database/sql"
	"html/template"
	"calculadora/backend/service"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("../frontend/*.html"))
}

type Controller struct {
	operationService service.OperationService
	db *sql.DB
}


func NewController(operationService service.OperationService, db *sql.DB) *Controller {
	return &Controller{
		operationService: operationService,
		db: db,
	}
}

