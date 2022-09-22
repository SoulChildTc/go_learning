package main

import "fmt"

func main() {
	var operator = "up"
	switch operator {
		case "start","up":
			fmt.Println("启动服务")
		case "stop":
			fmt.Println("停止服务")
		case "restart":
			fmt.Println("重启服务")
		case "reload":
			fmt.Println("重新加载")
		default:
			fmt.Println("Usage: start|stop|restart|reload")
	}


	// 穿透
	switch 2 {
	case 1:
		println("=1")
	case 2:
		println("=2")
		fallthrough //继续执行下一个case
	case 666:
		println("穿透过来吧")
	}
}
