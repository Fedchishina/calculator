package service

import (
	"calculator/application/entity"
	"calculator/application/stack"
	"strings"
)

func ToPostfix(expr entity.InfixExpr) entity.PostfixExpr {
	var str, postfixExpr string
	var st stack.Stack
	operationPriority := GetOperationPriority()
	symbols := strings.Split(string(expr), "")

	for i := 0; i < len(symbols); i++ {
		switch {
		case IsDigit(symbols[i]):
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

	return entity.PostfixExpr(postfixExpr)
}
