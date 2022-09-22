package constant_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota // 向左位移一位,高位丢弃，低位补0
	Writeable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}


func TestConstantTry1(t *testing.T) {
	a := 6
	/*
	6	0110
	1	0001
	r	0110

	6	0110
	2	0010
	r	0010

	6	0110
	4	0100
	r	0100
	*/
	t.Log(Readable, Writeable, Executable)
	t.Log(a&Readable, a&Writeable, a&Executable)
	t.Log(a&Readable==Readable, a&Writeable==Writeable, a&Executable==Executable)
}