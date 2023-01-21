package validator

import "fmt"

type Validator struct {
	Errors []string
}

func New() *Validator {
	return &Validator{Errors: nil}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(message string) {
	v.Errors = append(v.Errors, message)
}

func (v *Validator) Check(ok bool, message string) {
	if !ok {
		v.AddError(message)
	}
}

func (v *Validator) WriteErrors() {
	for _, err := range v.Errors {
		fmt.Println(err)
	}
}
