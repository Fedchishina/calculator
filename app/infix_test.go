package app

import (
	"testing"
)

func TestInfixExpr_Validate(t *testing.T) {
	tests := []struct {
		name    string
		i       Infix
		wantErr bool
	}{
		{"empty input data", Infix(""), true},
		{"wrong symbols", Infix("sss"), true},
		{"check balance", Infix("(1+1))"), true},
		{"success", Infix("(1+1)"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
