package main

import (
	"fmt"
)

// Human 定义结构体
type Human struct {
	name string
	age  int
}

func (h *Human) eat() {
	fmt.Println(h.name, "吃")
}

func (h *Human) walk() {
	fmt.Println(h.name, "走路")
}

// Hero 定义英雄结构体,继承人类结构体
type Hero struct {
	Human        // 继承Human
	skill string // 自有属性-技能
}

// 重写父类方法
func (h *Hero) eat() {
	fmt.Println(h.name, "英雄", "吃")
}

// 增加自有方法
func (h *Hero) fly() {
	fmt.Println(h.name, "在飞")
}

func main() {
	// 实例化一个人类结构体
	zs := Human{"张三", 12}
	zs.eat()
	zs.walk()

	// 实例化一个超人结构体
	superman := Hero{Human{"超人", 22}, "透视"}
	superman.eat()
	superman.walk()
	superman.fly()
}
