package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"calculadora/backend/model"
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


func NewController(operationService service.OperationService, *sql.DB) *Controller {
	return &Controller{
		operationService: operationService,
		db: db,
	}
}

