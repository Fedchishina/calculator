package main

import (
	"calculator/application/entity"
	"calculator/application/service"
	"calculator/application/validator"
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	ie := entity.InfixExpr(input)
	v := validator.New()

	if ie.Validate(v, input); !v.Valid() {
		v.WriteErrors()
		return
	}

	expr := service.ToPostfix(ie)
	fmt.Println(expr)
	fmt.Println(service.Calc(expr))
}
