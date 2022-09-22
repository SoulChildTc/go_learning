package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	t.Log(a)
	// 1 1 2 3 5 8 13 21 34
	for i := 0; i < 10; i++ {
		t.Log(b)
		tmp := a + b
		b = tmp
		a = b
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	tmp := a
	a = b
	b = tmp

	a,b = b,a
	t.Log(a,b)

}
