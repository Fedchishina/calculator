package entity

import (
	"calculator/application/validator"
)

type InfixExpr string

func (i *InfixExpr) Validate(v *validator.Validator, value string) {
	v.Check(value != "", "the input data is empty. Enter data please")
	v.Check(symbolsValid(v, value), "the input data is invalid")
}

func symbolsValid(v *validator.Validator, input string) bool {
	//validSymbols := getValidSymbols()
	//for _, s := range strings.Split(input, "") {
	//	if ! Contains(validSymbols, s) {
	//		v.AddError("Wrong symbol: " + s)
	//	}
	//}

	return v.Valid()
}

func getValidSymbols() []string {
	return []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		"+", "-", "/", "*",
		"(", ")",
		"^",
	}
}
