package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n1 int = 10
	s1 := strconv.FormatInt(int64(n1), 10)
	fmt.Printf("s1的类型是：%T, 值为%q \n",s1,s1)

	var n2 float32 = 3.1415
	// 参数1 要转换的值
	// 参数2 转换为10进制形式
	// 参数3 保留2为小数
	// 参数4 float64类型
	s2 := strconv.FormatFloat(float64(n2), 'f', 7, 64 )
	fmt.Printf("s2的类型是: %T, 值是%q \n", s2,s2)
}
