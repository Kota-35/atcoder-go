package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSum(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{10, 1},
		{11, 2},
		{1234, 10},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, calculateSum(tt.num))
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		n    int
		a    int
		b    int
		want int
	}{
		{20, 2, 5, 84},
		{10, 1, 2, 13},
		{100, 4, 16, 4554},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, solve(tt.n, tt.a, tt.b))
	}
}
