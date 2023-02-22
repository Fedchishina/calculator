package main

import (
	"testing"
)

func TestInfixExpr_Validate(t *testing.T) {
	tests := []struct {
		name    string
		i       InfixExpr
		wantErr bool
	}{
		{"empty input data", InfixExpr(""), true},
		{"wrong symbols", InfixExpr("sss"), true},
		{"check balance", InfixExpr("(1+1))"), true},
		{"success", InfixExpr("(1+1)"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
