package app

import (
	"strconv"
	"strings"
)

const (
	openingBracket = "("
	closingBracket = ")"
)

func isDigit(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}

func isOperator(s string) bool {
	return strings.Contains("+-/*^", s)
}

func findWholeNumber(expr string, pos int) (string, int) {
	var strNumber string
	exprArray := strings.Split(expr, "")
	for ; pos < len(exprArray); pos++ {
		s := exprArray[pos]
		if isDigit(s) {
			strNumber += s
		} else {
			pos--
			break
		}
	}

	return strNumber, pos
}
