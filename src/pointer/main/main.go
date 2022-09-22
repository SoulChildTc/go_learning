package main

import "fmt"

func test1()  {
	var num int = 10
	fmt.Println(num)

	var ptr *int = &num
	*ptr = 20
	fmt.Println(num)
}


func main() {
	/*
		所有基本数据类型都有对应的指针类型,前面加*就可以表示这个类型的指针类型，例如*int代表int的指针类型, *string代表string的指针类型
	 */

	var age int = 18
	// 通过&age获取变量内存地址, &age的类型是*int，int指针类型
	fmt.Println(&age)

	// 定义一个指针变量ptr，类型是：int类型的指针
	var ptr *int = &age
	fmt.Println(ptr)

	// 通过*ptr获取指针对应的值
	fmt.Println(*ptr,*&age)

	// 通过指针修改值
	test1()
}
