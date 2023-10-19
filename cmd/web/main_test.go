package main

import (
	"testing"
)

func Divide(a int, b int) int {
	var res int
	res = a / b
	return res

}

func TestDiv(t *testing.T) {
	aa := Divide(5, 5)
	t.Log("ok")
	_ = aa
}
