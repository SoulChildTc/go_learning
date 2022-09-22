package main

import "fmt"


func sum(x int, y int) (int, string) {
	return x + y, "test"
}

func exchangeNum(num1 int, num2 int) (int, int) {
	num1,num2 = num2,num1
	return num1,num2
}


func exchangeNum2(num1 *int, num2 *int) {
	// num1是int指针类型，里面存的是内存地址
	// 通过*num1可以表示num1中存的内存地址
	println(num1, *num1)
	*num1,*num2 = *num2,*num1
}


// args... int 可以传入0个，1个，N个int类型的参数
func test(args... int)  {
	for i:=0;i<len(args);i++ {
		fmt.Print(args[i])
	}
	fmt.Println()
}

func testR(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

func main() {
	// sum
	println("求和------------------------------------------------")
	res, _ := sum(1, 2)
	println(res)

	// 值交换
	println("值交换------------------------------------------------")
	num1 := 10
	num2 := 20
	println(exchangeNum(num1, num2))
	exchangeNum2(&num1, &num2)
	println(num1, num2)

	// 可变参数
	println("可变参数------------------------------------------------" )
	test()
	test(1)
	test(1,2,3,4,5)

	// 自动return
	print("自动return------------------------------------------------" )
	println(testR(30, 2))
}
