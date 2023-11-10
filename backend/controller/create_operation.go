package controller

import (
	"log"
	"net/http"
	"strconv"
	"calculadora/backend/domain"
)

func (c *Controller) CreateOperation(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		tmpl.ExecuteTemplate(w, "index.html", nil)
		return
	}

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
		return
	}

	tmpl.ExecuteTemplate(w, "index.html", struct { Result float64 } { Result: operation.Result })
}
