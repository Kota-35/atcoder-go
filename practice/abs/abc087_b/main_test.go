package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		c    int
		x    int
		want int
	}{
		{2, 2, 2, 100, 2},
		{5, 1, 0, 150, 0},
		{30, 40, 50, 6000, 213},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, solve(tt.a, tt.b, tt.c, tt.x))
	}
}
