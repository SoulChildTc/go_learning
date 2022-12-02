package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		goto可以将代码跳到指定的位置执行，注意goto执行完成会接着走goto下面的代码,知道没有代码或退出为止。
		goto一般会配合if使用
	*/

	goto label2
	fmt.Println("what")

label1:
	fmt.Println("Hello Golang1, stop")
	os.Exit(0) // 不推出会无限循环执行
label2:
	fmt.Println("Hello Golang2, goto 5")
	goto label5
label3:
	fmt.Println("Hello Golang3, goto 4")
	goto label4
label4:
	fmt.Println("Hello Golang4, goto 6")
	goto label6
label5:
	fmt.Println("Hello Golang5, goto 3")
	goto label3
label6:
	fmt.Println("Hello Golang6, goto 1")
	goto label1
}
