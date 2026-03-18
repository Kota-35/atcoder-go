package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountTwos(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{1, 0},
		{2, 1},
		{3, 0},
		{4, 2},
		{6, 1},
		{12, 2},
		{16, 4},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, countTwos(tt.num))
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		as   []int
		want int
	}{
		{[]int{8, 12, 40}, 2},
		{[]int{5, 6, 8, 10}, 0},
		{[]int{382253568, 723152896, 37802240, 379425024, 404894720, 471526144}, 8},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, solve(tt.as))
	}
}
