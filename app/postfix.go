package app

import (
	"github.com/Fedchishina/calculator/app/stack"
	"strings"
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
			{
				str, i = getStringNumber(string(expr), i)
				postfixExpr += str + " "
			}
		case symbols[i] == "(":
			st.Push("(")
		case symbols[i] == ")":
			{
				for st.Count() > 0 && st.Top() != "(" {
					postfixExpr += st.Top() + " "
					st.Pop()
				}
				st.Pop()
			}
		default:
			{
				op := symbols[i]
				for st.Count() > 0 && operationPriority[st.Top()] >= operationPriority[op] {
					postfixExpr += st.Top() + " "
					st.Pop()
				}
				st.Push(op)
			}
		}
	}

	for j := 0; j < st.Count(); j++ {
		postfixExpr += st.Top() + " "
		st.Pop()
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

func validOperators() []string {
	return []string{
		"+", "-", "/", "*", "^",
	}
}
