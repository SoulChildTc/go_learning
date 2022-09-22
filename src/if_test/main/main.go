package main

import "fmt"

func main() {
	var num int
	fmt.Print("输入你的数字: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println(err)
		return
	}

	if num > 100 {
		fmt.Println("你赢了")
	}else {
		fmt.Println("你输了")
	}


	// 新玩法
	if count := 10; count < 100 {
		fmt.Println("10小于100")
	}
}
