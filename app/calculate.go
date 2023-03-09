package app

import (
	"fmt"
	"math"
	"unicode"

	"github.com/Fedchishina/calculator/app/stack"
)

func Calculate(expr Postfix) (string, error) {
	var numbers stack.Stack
	var number string
	var first, second float64
	var pos int

	for i, val := range expr {
		switch {
		case i < pos:
			continue
		case unicode.IsDigit(val):
			number, pos = findWholeNumber(string(expr), pos)
			numbers.Push(number)
		case isOperator(val):
			second = numbers.PopNumber()
			first = numbers.PopNumber()
			ex := execute(val, first, second)
			numbers.Push(fmt.Sprintf("%f", ex))
		}
		pos++
	}

	return numbers.Pop(), nil
}

func execute(op int32, first float64, second float64) float64 {
	switch op {
	case '+':
		return first + second
	case '-':
		return first - second
	case '*':
		return first * second
	case '/':
		return first * second
	case '^':
		return math.Pow(first, second)
	}
	return 0
}
