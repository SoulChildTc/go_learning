package main

import (
	"fmt"
)

// map
func main() {
	// 定义map 方式1
	fmt.Println("方式1================================================================================")
	// var m1 map[int]string //可省略
	m1 := make(map[int]string,3) //可以存10对映射关系
	m1[2001] = "张三"
	m1[2002] = "李四"
	m1[2003] = "王五"
	fmt.Println("长度", len(m1))
	fmt.Println("m1: ",m1)
	m1[2004] = "王六" // 证明map是可变长的
	fmt.Println("长度", len(m1))
	fmt.Println("m1: ",m1)

	fmt.Println("方式2================================================================================")
	// 定义map 方式2,不指定大小
	m2 := make(map[int]string)
	m2[2001] = "soulchild"
	fmt.Println("m2: ",m2)

	fmt.Println("方式3================================================================================")
	// 定义map 方式2,直接赋值定义
	m3 := map[int]string{
		2001:"soul",
		2002:"child",
	}
	fmt.Println("m3: ",m3)
	m3[2003] = "by"
	fmt.Println("m3: ",m3)

	fmt.Println("删除key================================================================================")
	delete(m3,2003)
	fmt.Println("删除了m3[2003]: ",m3)

	fmt.Println("清空map================================================================================")
	m1 = make(map[int]string) // make生成一个新的，原来的成为垃圾，通过gc回收
	fmt.Println("清空map m1: ",m1)

	// 通过循环删除
	for k,_ := range m2{
		delete(m2,k)
	}
	fmt.Println("清空map m2: ",m2)

	fmt.Println("查找key================================================================================")
	v,flag := m3[2001]
	fmt.Println(v)
	fmt.Println(flag)
	v1,flag1 := m3[2003]
	fmt.Println(v1)
	fmt.Println(flag1)

	fmt.Println("嵌套map================================================================================")
	m4 := make(map[string]map[int]string) // 含义: map类型key是string类型，value是map类型，子map的key类型为int，value为string
	m4["班级1"] = map[int]string{} // 创建一个子map
	m4["班级1"][2001] = "张三"
	m4["班级1"][2002] = "李四"
	m4["班级1"][2003] = "王五"
	m4["班级2"] = map[int]string{} // 创建一个子map
	m4["班级2"][2001] = "soul"
	m4["班级2"][2002] = "child"
	fmt.Println(m4)
	m5 := make(map[string][]int)
	m5["key1"] = []int{1,2,3,4}
	fmt.Println(m5)

}
