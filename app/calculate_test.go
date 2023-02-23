package app

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	type args struct {
		expr Postfix
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"example 1", args{Postfix("2 3 + 4 * ")}, "20.000000", false},
		{"example 2", args{Postfix("2 3 4 * + ")}, "14.000000", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.args.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
