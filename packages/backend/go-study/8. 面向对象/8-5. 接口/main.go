package main

import (
	"fmt"
)

// 首字母大小写 同接口，函数同一个意义
// 实现一个接口要求，就是将接口内的方法都实现了 才叫实现了这个接口
type PersonInterface interface {
	// 接口内定义的方法，不需要方法体
	SayHello()
	SayGoodBye()
}

type Custom interface {
	sayYT()
	// CustomOne  //  Custom 接口 继承 CustomOne 接口
}

type CustomOne interface {
	SayJk()
}

type Person struct {
	Age  int
	Name string
}

// 此时，Person 结构体就 实现了  PersonInterface接口
// 因为SayGoodBye() 和 SayHello() 方法都定义在了Person 结构体中了
func (p Person) SayGoodBye() {
	fmt.Println("SayGoodBye")
}

func (p Person) SayHello() {
	fmt.Println("SayHello")
}

// Person 结构体 同时实现了 Custom 接口
func (p Person) sayYT() {
	fmt.Println("SayHello")
}

// 第一第二个结构体， 也实现上述的接口
type Student struct {
	Age  int
	Name string
}

func (s Student) SayHello() {
	fmt.Println("Student SayHello")
}
func (s Student) SayGoodBye() {
	fmt.Println("Student SayGoodBye")
}

type Teacher struct {
}

func (s Teacher) SayGoodBye() {
	fmt.Println("Student SayGoodBye")
}

func main() {
	// Student 和 Person 都实现了 PersonInterface 接口，所以可将实例传入到greet 函数中
	greet(Person{18, "zhanghaojie"})
	greet(Student{18, "zhanghaojie"})

	// Teacher 结构体没有实现 PersonInterface 接口，所以报错
	// greet(Teacher{})

	// 传任意类型 都不报错
	hello("a")
	hello(1111)
	hello(22.3)
	hello(false)
	hello([]int{1, 3, 4})
	hello(map[string]int{"a": 1, "2": 3, "3": 4})
}

// 只要实现了这个接口的结构体，都是可以传入这个函数的
func greet(s PersonInterface) {
	s.SayGoodBye()
}

// 此时任意类型的 参数都可以传入到这个函数
func hello(a interface{}) {
	fmt.Println(a)
}
