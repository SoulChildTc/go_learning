package main

import "fmt"

func main() {
	fmt.Print("请输入你的年龄: ")
	var age int
	fmt.Scanln(&age)

	fmt.Print("请输入你的姓名: ")
	var name string
	fmt.Scanln(&name)

	fmt.Printf("姓名: %v, 年龄: %v \n", name, age)

	//////////////自定义格式////////////////////
	fmt.Print(`请输入你的信息,格式"年龄-姓名": `)
	fmt.Scanf("%d-%s", &age,&name)
	fmt.Printf("姓名: %v, 年龄: %v \n", name, age)

}
