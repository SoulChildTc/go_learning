package main

import "fmt"

// 定义结构体
type Hero struct {
	name  string
	age   int
	level int
}

// 为结构体添加方法,定义方法的语法如下
// 可以看到在方法名前面多了一个(h *Hero),其中h可以理解为其他语言的this、self, *Hero是类型。这样写了以后代表将setName绑定到Hero结构体上。
// 当我们通过xx.setName调用这个方法的时候,他会把xx传给h，所以这边Hero我们使用的是指针类型,如果不是指针类型修改不会生效,因为go默认是值拷贝
func (h *Hero) setName(name string) {
	h.name = name
}

// 读取操作其实也可以不用指针类型
func (h *Hero) show() {
	fmt.Println(h.name)
	fmt.Println(h.age)
	fmt.Println(h.level)
}

func main() {
	// 实例化一个结构体
	var riven Hero
	// 调用setName方法
	riven.setName("锐雯")
	fmt.Println(riven)
	riven.show()
}
