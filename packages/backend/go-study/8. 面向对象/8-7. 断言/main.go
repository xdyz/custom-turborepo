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
	var a interface{}

	a = "hello"

	// 1. 通过断言获取类型
	s, ok := a.(interface{}) // 断言正确 s 为 a 的值
	s1, ok1 := a.(bool)      // 断言错误，s1 为 false

	fmt.Printf("ok: %v\n", ok)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("ok1: %v\n", ok1)
	fmt.Printf("s1: %v\n", s1)
}

func greet(s PersonInterface) {
	// 断言接口是不是一个Student 类型结构体
	if val, ok := s.(Student); ok {
		val.SayGoodBye()
	}

	// type 为关键字，必须这么写，才能用于判断类型
	switch s.(type) {
	case Student:
		s1, _ := s.(Student)
		s1.SayGoodBye()
	case Person:
		p, _ := s.(Person)
		p.SayGoodBye()
	}
}
