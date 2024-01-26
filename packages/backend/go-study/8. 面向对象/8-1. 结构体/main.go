package main

import (
	"fmt"
)

// 定义一个结构体
type Teacher struct {
	Name    string
	Age     int
	Subject string
	address string // 小写的属性，外部无法访问
}

// 这个方法就属于 Teacher 这个结构体
func (t Teacher) SayHello() {
	fmt.Println("Hello, I am", t.Name)
}

func main() {

	// 实例化一个结构体
	teacher := Teacher{
		Name:    "张三",
		Age:     20,
		Subject: "语文",
		address: "北京",
	}

	teacher.SayHello()

	// 实例化一个结构体指针
	teacherPtr := &teacher
	(*teacherPtr).Name = "李四"

	var t1 *Teacher = &teacher
	(*t1).address = "上海"
	t1.Name = "李四" // 这种方式，编译器底层自动转为 (*t1).Name = "李四"

	// 也可以使用 new() 函数创建一个指针
	var t2 *Teacher = new(Teacher)
	(*t2).Name = "李四"
	t2.Age = 21 // 这种方式，编译器底层自动转为 (*t2).Age = 21

	// 上面实例化指针类型 比较麻烦 还需要一个一个赋值，还需要使用new方法
	// 下面可以简化，直接使用结构体的初始化语法
	var t3 *Teacher = &Teacher{
		Name:    "李四",
		Age:     21,
		Subject: "数学",
		address: "深圳",
	}

	fmt.Println(t3)

	fmt.Println(teacher)

}
