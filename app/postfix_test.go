package app

import (
	"testing"
)

func TestToPostfix(t *testing.T) {
	type args struct {
		expr Infix
	}
	tests := []struct {
		name string
		args args
		want Postfix
	}{
		{"example 1", args{Infix("(2+3)*4")}, Postfix("2 3 + 4 * ")},
		{"example 2", args{Infix("2+(3*4)")}, Postfix("2 3 4 * + ")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPostfix(tt.args.expr); got != tt.want {
				t.Errorf("ToPostfix() = %v, want %v", got, tt.want)
			}
		})
	}
}
