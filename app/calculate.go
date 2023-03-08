package app

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/Fedchishina/calculator/app/stack"
)

func Calculate(expr Postfix) (string, error) {
	var numbers stack.Stack
	var number string
	var first, second float64
	var err error

	symbols := strings.Split(string(expr), "")
	for i := 0; i < len(symbols); i++ {
		if isDigit(symbols[i]) {
			number, i = findWholeNumber(string(expr), i)
			numbers.Push(number)
		}
		if isOperator(symbols[i]) {
			second, err = strconv.ParseFloat(numbers.Pop(), 64)
			if err != nil {
				log.Fatalf("error during parsing %v to float : %v", first, err)
			}
			first, err = strconv.ParseFloat(numbers.Pop(), 64)
			if err != nil {
				log.Fatalf("error during parsing %v to float : %v", first, err)
			}

			ex := execute(symbols[i], first, second)
			numbers.Push(fmt.Sprintf("%f", ex))
		}
	}

	return numbers.Pop(), nil
}

func execute(op string, first float64, second float64) float64 {
	switch op {
	case "+":
		return first + second
	case "-":
		return first - second
	case "*":
		return first * second
	case "/":
		return first * second
	case "^":
		return math.Pow(first, second)
	}
	return 0
}
