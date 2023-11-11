package controller

import (
	"calculadora/backend/domain"
	"log"
	"net/http"
	"strconv"
)

func (c *Controller) CreateOperation(w http.ResponseWriter, r *http.Request){
	var operations []domain.Operation
	operations, err := c.operationService.GetOperations(c.db)
	log.Print(operations)

	if err != nil {
		log.Fatal(err)
	}

	if r.Method == "POST" {
		result := c.cleanData(r)

		tmpl.ExecuteTemplate(w, "index.html", struct { 
			Result float64
			Operations []domain.Operation
		} { 
			Result: result, 
			Operations: operations,
		})
	} else {
		tmpl.ExecuteTemplate(w, "index.html", struct { 
			Operations []domain.Operation
		} { 
			Operations: operations,
		})
	} 
}

func (c *Controller) cleanData(r *http.Request) float64 {
	r.ParseForm()
	input := r.PostFormValue("input")
	slicedString := sliceString(input)

	numberint, _ := strconv.Atoi(slicedString[0])
	number1 := float64(numberint)

	numberint, _ = strconv.Atoi(slicedString[1])
	number2 := float64(numberint)

	operator := slicedString[2]

	var result float64
	if operator == "+" {
		result = number1 + number2	
	} else if operator == "-" {
		result = number1 - number2
	} else if operator == "*" {
		result = number1 * number2
	} else if operator == "/" {
		result = number1 / number2
	}

	operationInfo := domain.Operation{
		Number1: number1,
		Number2: number2,
		Operator: operator,
		Result: result,
	}

	operation, err := c.operationService.CreateOperation(operationInfo, c.db)

	if err != nil {
		log.Fatal(err)
	}
	 
	return operation.Result
}
