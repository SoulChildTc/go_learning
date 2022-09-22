package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		goto可以将代码跳到指定的位置执行，注意goto执行完成会接着走goto下面的代码。
		所以这里相当于两条线在执行代码，一条是指定位置的代码会一直执行到没有代码为止，一条是goto下面的代码执行到没有代码为止
		goto一般会配合if使用
	*/

	goto label2

label1:
	fmt.Println("Hello Golang1, stop")
	os.Exit(0)
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
