package service

import (
	"strconv"
	"strings"
)

func GetOperationPriority() map[string]int {
	return map[string]int{
		"(": 0,
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
	}
}

func getValidOperators() []string {
	return []string{
		"+", "-", "/", "*", "^",
	}
}

func GetStringNumber(expr string, pos int) (string, int) {
	var strNumber string
	exprArray := strings.Split(expr, "")
	for ; pos < len(exprArray); pos++ {
		s := exprArray[pos]
		if IsDigit(s) {
			strNumber += s
		} else {
			pos--
			break
		}
	}

	return strNumber, pos
}

func IsDigit(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
