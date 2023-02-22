package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	ie := InfixExpr(input)

	if err := ie.Validate(); err != nil {
		fmt.Println(err)
		return
	}

	expr := ToPostfix(ie)
	fmt.Println(expr)
	fmt.Println(Calculate(expr))
}
