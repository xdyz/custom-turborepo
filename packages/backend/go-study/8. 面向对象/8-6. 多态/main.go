package main

import "fmt"

// 首字母大小写 同接口，函数同一个意义
// 实现一个接口要求，就是将接口内的方法都实现了 才叫实现了这个接口
type PersonInterface interface {
	// 接口内定义的方法，不需要方法体
	SayHello()
	SayGoodBye()
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

func main() {

	// 多态数组，其实就是不同的结构体实现了 这个接口，其组合起来就成了一个多态数组
	p := [2]PersonInterface{
		Person{Age: 10, Name: "tom"},
		Student{Age: 10, Name: "jack"},
	}

	for _, v := range p {
		greet(v)
	}
}

// 此时下面的函数就是一种多态，只要是实现了这个接口的结构体，都可以传入这个函数
// 只要实现了这个接口的结构体，都是可以传入这个函数的
func greet(s PersonInterface) {
	// 断言接口是不是一个Student 类型结构体
	if val, ok := s.(Student); ok {
		val.SayGoodBye()
	}
}
