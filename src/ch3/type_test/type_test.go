package type_test

import (
	"math"
	"reflect"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T)  {
	var a int32 = 1
	var b int64
	var c MyInt
	b = int64(a)
	c = MyInt(a)
	t.Log(a, b, c)
}

func TestImplicit2(t *testing.T)  {
	a := 1.2
	t.Log(reflect.TypeOf(a))
	t.Log(math.MinInt64)
	t.Log(math.MaxInt64)
}
