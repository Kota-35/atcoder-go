package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		as    []int
		wantt int
	}{
		{[]int{3, 1}, 2},
		{[]int{2, 7, 4}, 5},
		{[]int{20, 18, 2, 18}, 18},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.wantt, solve(tt.as))
	}
}
