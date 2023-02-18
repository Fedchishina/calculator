package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	ie := InfixExpr(input)
	v := New()

	if ie.Validate(v, input); !v.Valid() {
		v.WriteErrors()
		return
	}

	expr := ToPostfix(ie)
	fmt.Println(expr)
	fmt.Println(Calculate(expr))
}
