package main

import (
	"strings"

	"github.com/Fedchishina/calculator/stack"
)

type PostfixExpr string

func ToPostfix(expr InfixExpr) PostfixExpr {
	var str, postfixExpr string
	var st stack.Stack
	operationPriority := operatorsPriority()
	symbols := strings.Split(string(expr), "")

	for i := 0; i < len(symbols); i++ {
		switch {
		case isDigit(symbols[i]):
			{
				str, i = GetStringNumber(string(expr), i)
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

	return PostfixExpr(postfixExpr)
}