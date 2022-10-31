package test

import (
	"testing"

	"go-tdd/src"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := src.Dollar{}
	five.New(5)
	product := five.Times(2)
	assert.Equal(t, 10, five.Amount)
	product = five.Times(3)
	assert.Equal(t, 15, five.Amount)
}
