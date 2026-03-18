package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		ds   []int
		want int
	}{
		{[]int{10, 8, 8, 6}, 3},
		{[]int{15, 15, 15}, 1},
		{[]int{50, 30, 50, 100, 50, 80, 30}, 4},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, solve(tt.ds))
	}
}
