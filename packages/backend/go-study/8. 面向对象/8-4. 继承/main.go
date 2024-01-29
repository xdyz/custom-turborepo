package main

import "fmt"

type person struct {
	Name string
	Age  int
	age1 int
}

type Others struct {
	hasFmt bool
}

func (p person) GetAge1() int {
	return p.age1
}

func (p *person) SetAge1(age int) {
	p.age1 = age
}

// 嵌套一个匿名结构体，实现继承
type Student struct {
	person
	Others
	Grade int
	Name  string // 这里会覆盖Person中的Name属性
}

func (s Student) GetGrade() int {
	return s.Grade
}

// 可以给嵌套的结构体起名，通过名称访问结构体的属性和方法
type Teacher struct {
	person
	Subject string
	address string
}

func (t Teacher) GetAddress() string {
	return t.address
}

// 可以给嵌套的结构体起名，通过名称访问结构体的属性和方法
type NewTeacher struct {
	*person
	p person // 这种带别称是 组合关系 不能算继承关系
}

func main() {
	s := Student{
		person: person{
			Name: "zhangsan",
			Age:  18,
		},
		Grade: 100,
		Others: Others{
			hasFmt: true,
		},
	}

	fmt.Println(s.Age)
	fmt.Println(s.Name)
	fmt.Println(s.Grade)
	fmt.Println(s.GetAge1())
	fmt.Println(s.Others.hasFmt) // 上下两种访问当时都是可以的
	fmt.Println(s.hasFmt)

	t := Teacher{
		person: person{
			Name: "lisi",
			Age:  20,
		},
		Subject: "math",
		address: "beijing",
	}

	// 此时访问name 就需要通过t.person.Name 来访问
	fmt.Println(t.person.Name)
	fmt.Println(t.Subject)
	fmt.Println(t.person.GetAge1())

	t.person.SetAge1(20) // 此时 t 会被自动转为 person 的指针

	t1 := NewTeacher{&person{"wangwu", 20, 33}, person{"wangwu", 20, 33}}

	// 这两种都可以
	fmt.Println(t1.person.Age) // 自动访问转换了
	fmt.Println((*t1.person).Age)
}
