package service

import (
	"calculator/application/entity"
	"calculator/application/stack"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Calc(expr entity.PostfixExpr) string {
	var numbers stack.Stack
	var number string
	var first, second float64
	operators := getValidOperators()
	symbols := strings.Split(string(expr), "")

	for i := 0; i < len(symbols); i++ {
		if IsDigit(symbols[i]) {
			number, i = GetStringNumber(string(expr), i)
			numbers.Push(number)
		} else if Contains(operators, symbols[i]) {
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
	default:
		return 0
	}
}
