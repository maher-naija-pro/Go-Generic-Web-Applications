package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Divide(a int, b int) int {
	var res int
	res = a / b
	return res

}

func TestDiv(t *testing.T) {
	assert.False(t, false, "its fals")
	aa := Divide(5, 5)
	_ = aa
}
