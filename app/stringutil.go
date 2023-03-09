package app

import (
	"unicode"
)

func isOperator(r int32) bool {
	switch r {
	case '+', '-', '*', '/', '^':
		return true
	}
	return false
}

func findWholeNumber(expr string, pos int) (string, int) {
	var strNumber []rune
	for _, val := range expr[pos:] {
		if unicode.IsDigit(val) {
			strNumber = append(strNumber, val)
			pos++
		} else {
			break
		}
	}
	return string(strNumber), pos
}
