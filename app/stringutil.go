package app

import (
	"strconv"
	"strings"
	"unicode"
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
	var strNumber []rune
	for _, val := range expr[pos:] {
		if unicode.IsDigit(val) {
			strNumber = append(strNumber, val)
			pos++
		} else {
			pos--
			break
		}
	}
	return string(strNumber), pos
}
