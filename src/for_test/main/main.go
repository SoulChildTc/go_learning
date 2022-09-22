package main

import "fmt"

func main() {
	// for常规用法
	var res int
	for i := 1; i <= 5; i++ {
		res += i
		if res == 10 { // 提前终止循环
			break
		}
	}
	fmt.Println(res)

	// for ... range
	// 目前仅发现能代替for i:=0;i<len(arr);i++{}这种，并不像python那样可以range(1,20)
	txt := "hello world"
	for i, value := range txt {
		fmt.Printf("索引%v,值%c \n", i, value)
	}

	// 循环加标签, 可以指定break哪一层循环
outer:
	for i := 0; i < 5; i++ {
	inner:
		for j := 0; j < 5; j++ {
			if i == 2 {
				break inner
			} else if i == 4 {
				break outer
			}
			fmt.Println(i, j)
		}
	}

	// continue, 结束本次循环，继续下次循环
	// 输出1-100被6整除的数
	for i := 0; i < 100; i++ {
		if i%6 != 0 {
			continue
		}
		fmt.Println(i)
	}

	// continue加标签
outer2:
	for i := 0; i < 5; i++ {
		fmt.Printf("\n%v", i)
	inner2:
		for j := 0; j < 5; j++ {
			if i == 2 {
				continue outer2 //结束本次outer2循环，继续outer2下次循环
			} else if j == 0 {
				continue inner2 //结束本次inner2循环，继续inner2下次循环
			}
			fmt.Print(j)
		}
	}

	// 死循环
	//for {
	//	fmt.Println("\n死循环，ctrl+c吧")
	//	fmt.Scanln()
	//}
}
