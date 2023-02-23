package app

import (
	"strconv"
	"strings"
)

func isDigit(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func getStringNumber(expr string, pos int) (string, int) {
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
