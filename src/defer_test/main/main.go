package main

import "fmt"

func main() {
	// defer会将后面的内容压入栈(先进后出),在函数执行到最后时,会将栈的内容进行依次执行
	//defer fmt.Println("main end1")  // 第一个进栈，最后执行
	//defer fmt.Println("main end2")  // 第二个进栈，倒数第二个执行
	//
	//fmt.Println("main::hello 1")
	//fmt.Println("main::hello 2")

	defer fun1()
	defer fun2()
	defer fun3()
	testDeferAndReturn() //return先执行、defer最后执行
}

func fun1() {
	fmt.Println("1")
}

func fun2() {
	fmt.Println("2")
}

func fun3() {
	fmt.Println("3")
}

func testDeferAndReturn() int {
	// 测试defer和return谁先执行
	defer deferFun()
	return returnFun()
}

func deferFun() int {
	fmt.Println("我是defer")
	return 0
}

func returnFun() int {
	fmt.Println("我是return")
	return 0
}
