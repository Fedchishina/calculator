package app

import (
	"strings"

	"github.com/Fedchishina/calculator/app/stack"
)

type Postfix string

func ToPostfix(expr Infix) Postfix {
	var str, postfixExpr string
	var st stack.Stack
	operationPriority := operatorsPriority()
	symbols := strings.Split(string(expr), "")

	for i := 0; i < len(symbols); i++ {
		switch {
		case isDigit(symbols[i]):
			str, i = findWholeNumber(string(expr), i)
			postfixExpr += str + " "
		case symbols[i] == openingBracket:
			st.Push(openingBracket)
		case symbols[i] == closingBracket:
			for st.Count() > 0 && st.Top() != openingBracket {
				postfixExpr += st.Pop() + " "
			}
			st.Pop()
		default:
			op := symbols[i]
			for st.Count() > 0 && operationPriority[st.Top()] >= operationPriority[op] {
				postfixExpr += st.Pop() + " "
			}
			st.Push(op)
		}
	}

	for st.Count() > 0 {
		postfixExpr += st.Pop() + " "
	}

	return Postfix(postfixExpr)
}

func operatorsPriority() map[string]int {
	return map[string]int{
		"(": 0,
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
	}
}
