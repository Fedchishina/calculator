package main

import (
	"fmt"
	"github.com/Fedchishina/calculator/app"
)

func main() {
	var input string
	fmt.Scan(&input)

	ie := app.Infix(input)

	if err := ie.Validate(); err != nil {
		fmt.Println(err)
		return
	}

	expr := app.ToPostfix(ie)
	fmt.Println(expr)
	fmt.Println(app.Calculate(expr))
}
