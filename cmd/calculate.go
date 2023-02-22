package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Fedchishina/calculator/stack"
)

func Calculate(expr PostfixExpr) string {
	var numbers stack.Stack
	var number string
	var first, second float64
	operators := validOperators()
	symbols := strings.Split(string(expr), "")

	for i := 0; i < len(symbols); i++ {
		if isDigit(symbols[i]) {
			number, i = getStringNumber(string(expr), i)
			numbers.Push(number)
		} else if contains(operators, symbols[i]) {
			first, _ = strconv.ParseFloat(numbers.Top(), 64)
			numbers.Pop()
			second, _ = strconv.ParseFloat(numbers.Top(), 64)
			numbers.Pop()
			ex := execute(symbols[i], first, second)

			numbers.Push(fmt.Sprintf("%f", ex))
		}
	}

	return numbers.Top()
}

func execute(op string, first float64, second float64) float64 {
	switch {
	case op == "+":
		return first + second
	case op == "-":
		return first - second
	case op == "*":
		return first * second
	case op == "/":
		return first * second
	case op == "^":
		return math.Pow(first, second)
	}
	return 0
}
