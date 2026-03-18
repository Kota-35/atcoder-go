package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"101", 2},
		{"000", 0},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, solve(tt.s))
	}
}
