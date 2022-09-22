package main

import "fmt"

func main() {
	// 简单数组使用
	var scores [5]int
	for i := 0; i < len(scores); i++ {
		fmt.Print("请输入学生成绩: ")
		//fmt.Scanln(&scores[i])
	}

	println("===========================================")
	for key, value := range scores {
		fmt.Printf("索引%v，成绩%v \n", key, value)
	}

	// 数组定义的4种方式
	// 第一种,完整写法
	var arr1 [3]int = [3]int{1,2,3}
	fmt.Println(arr1)
	// 第二种,类型推断
	var arr2 = [3]int{1,2,3}
	fmt.Println(arr2)
	// 第三种,不知道长度的
	var arr3 = [...]int{1,2,3,4,5}
	fmt.Println(arr3)
	// 第四种,指定索引定义对应的值
	var arr4 = [...]int{1:10,3:30,2:20,0:0}
	fmt.Println(arr4)
	println("二维数组===========================================")
	// 二维数组,arr5的长度为2,内部每个元素也是数组，长度为3
	var arr5 [2][3]int
	fmt.Println("arr5: ",arr5)
	// arr5的内存地址等于arr5[0]的地址,arr5[0]的内存地址等于arr5[0][0]的地址
	fmt.Printf("arr5的内存地址是%p, arr5[0][0]的内存地址是%p \n", &arr5, &arr5[0][0])
	// int占用8字节,所以内存地址的推断arr5[0][1] = arr5[0][0] + 8, arr5[0][2] = arr5[0][1] + 8。 注意这里是16进制的运算
	fmt.Printf("arr5[0][1]的内存地址是%p \n", &arr5[0][1])
	fmt.Printf("arr5[0][2]的内存地址是%p \n", &arr5[0][2])

	// arr[1]的地址等于arr5[0][2] + 8
	fmt.Printf("arr5[1]的内存地址是%p \n", &arr5[1])

	// 二维数组初始化
	var arr6 = [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(arr6)

}
