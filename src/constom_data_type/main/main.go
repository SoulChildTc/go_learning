package main

import "fmt"

type myInt int
type myFunc func()

func test0() {
	fmt.Println("test0")
}

func test(num1 int, num2 myInt, testFunc myFunc) {
	fmt.Println("test")
}

func main() {
	var a myInt = 2
	test(1, a, test0)
}
