package app

import (
	"unicode"

	"github.com/Fedchishina/calculator/app/stack"
)

type Postfix string

func ToPostfix(expr Infix) Postfix {
	var str, postfixExpr string
	var st stack.Stack
	var pos int

	operationPriority := operatorsPriority()
	for i, val := range expr {
		switch {
		case i < pos:
			continue
		case unicode.IsDigit(val):
			str, pos = findWholeNumber(string(expr), pos)
			postfixExpr += str + " "
			continue
		case val == '(':
			st.Push("(")
			pos++
			continue
		case val == ')':
			for st.Count() > 0 && st.Top() != "(" {
				postfixExpr += st.Pop() + " "
			}
			st.Pop()
			pos++
			continue
		}
		op := string(val)
		for st.Count() > 0 && operationPriority[st.Top()] >= operationPriority[op] {
			postfixExpr += st.Pop() + " "
		}
		st.Push(op)
		pos++
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
