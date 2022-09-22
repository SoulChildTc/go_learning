package main

import "fmt"

// 多态通过interface来实现

// AnimalIF interface更像是一个规范,要求了动物类必须包含的方法
type AnimalIF interface {
	sleep()
	GetColor() string // 获取动物的颜色,返回值是string
	GetType() string  // 获取动物的种类,返回值是string
}

// 定义类猫类
type cat struct {
	color   string
	variety string // 品种
}

// 实现动物接口要求的方法
func (c *cat) sleep() {
	fmt.Println("猫在睡觉")
}

func (c *cat) GetColor() string {
	fmt.Println("猫的颜色", c.color)
	return c.color
}

func (c *cat) GetType() string {
	return c.variety
}

// 定义狗类
type dog struct {
	color   string
	variety string
}

// 实现动物接口要求的方法
func (d *dog) sleep() {
	fmt.Println("狗在睡觉")
}

func (d *dog) GetColor() string {
	fmt.Println("狗的颜色", d.color)
	return d.color
}

func (d *dog) GetType() string {
	return d.variety
}

func showAnimal(animal AnimalIF) {
	animal.sleep()
	animal.GetColor()
	fmt.Println(animal.GetType())
}

func main() {
	// 上面通过一个接口实现了两个类

	// 实例化猫类
	var animal AnimalIF // 数据类型是AnimalIF,他是个接口类型
	/*
		这里是否使用取址符取决于，实现的方法中有没有使用指针接收器。
		如果有一个方法使用了指针接收器,那么就必须使用取址符。
		不管方法有没有使用指针接收器都可以加取址符
	*/
	animal = &cat{"白色", "加菲猫"}
	animal.sleep()
	animal.GetColor()
	fmt.Println(animal.GetType())

	// 实例化狗类
	animal = &dog{
		color:   "黄色",
		variety: "拉布拉多",
	}
	animal.sleep()
	animal.GetColor()
	fmt.Println(animal.GetType())

	// 通过函数去调用动物类的属性
	showAnimal(animal)
}
