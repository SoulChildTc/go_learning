package main

import "fmt"

// 切片
func main() {
	// 定义数组
	var intarr = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(intarr)
	fmt.Println("============================================================")

	//切片，与python一样前包后不包
	//var slice []int = intarr[1:3]  // []int可以省略
	var slice = intarr[1:3]
	slice2 := intarr[1:3]
	fmt.Println(slice)
	fmt.Println(slice2)
	// 容量是指这个数组最大能放多少个元素，长度是这个数组当前有多少个元素
	fmt.Println("intarr容量: ", cap(intarr))
	fmt.Println("slice容量: ", cap(slice)) // 因为是从1开始切的，所以容量等于原数组的容量-1，cap(intarr)-1
	// 底层指向的是同一个内存地址
	fmt.Println("intarr[1]的内存地址: ", &intarr[1])
	fmt.Println("slice[0]的内存地址: ", &slice[0])

	// 改动是互相影响的
	intarr[1] = 20
	fmt.Println(intarr, slice, slice2)
	slice[0] = 200
	fmt.Println(intarr, slice, slice2)

	fmt.Println("============================================================")
	// make定义切片. make会在底层创建一个数组,无法直接使用,slice3会引用他
	slice3 := make([]int,4,4)
	//slice3 := []int{0,0,0,0} // 等同make
	fmt.Println("slice3: ", slice3)
	fmt.Println("slice3的长度", len(slice3))
	fmt.Println("slice3的容量", cap(slice3))

	//此时，长度为4，容量为4
	//但是如果使用append()，切片长度会变为5 切片长度>容量，此时，切片的容量增大变为cap = len*2
	slice3 = append(slice3, 0)
	fmt.Println("slice3: ", slice3)
	fmt.Println("slice3的长度", len(slice3))
	fmt.Println("slice3的容量", cap(slice3))
	// 追加切片需要...
	slice3 = append(slice3,slice...)
	fmt.Println(slice3)

	fmt.Println("============================================================")
	// 值copy
	// 先定义一个长度为10的切片
	slice4 := make([]int,10)
	// 将slice3的每个元素按照对应的位置拷贝给slice4
	copy(slice4, slice3)
	fmt.Println(slice4)
	// copy后在修改slice4的元素，就不会影响slice3了
	slice4[5] = 500
	fmt.Println(slice3)
	fmt.Println(slice4)
}
