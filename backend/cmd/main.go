package main

import (
	"calculadora/backend/controller"
	"calculadora/backend/repository"
	"calculadora/backend/service"
	"embed"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main(){
	db, err := repository.NewDB()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	mux := mux.NewRouter()

	dao := repository.NewDAO()
	operatorService := service.NewOperationService(dao)
	c := controller.NewController(operatorService, db)

	var static embed.FS

	mux.Handle("frontend/", http.FileServer(http.FS(static)))
	mux.HandleFunc("/", c.CreateOperation)

	http.ListenAndServe(":8080", context.ClearHandler(mux))

}
