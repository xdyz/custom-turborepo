package main

import "fmt"

// 起别名的类型，可以给它添加方法
type A int

type Maps map[string]int

type Student struct {
	Name string
}

func main() {
	var a A = 10

	a.test()

	var b Maps = Maps{
		"1": 1,
	}

	b.test1()

	s := Student{
		Name: "1",
	}

	//  分别可以用值类型和指针类型来到用
	// s.test()
	(&s).test()                   // 虽然看起来是用指针，但是实际底层还是值类型
	fmt.Println("out = ", s.Name) // 就算传入的是指针，此时s 的值还是没有变

	// 上面反过来同样可以使用
	s2 := Student{
		Name: "2",
	}

	// 都是可以调用的
	s2.test2() // 虽然用的值传递，但是底层此时转为了指针类型
	fmt.Println(s2.Name)
	// (&s2).test2()
}

// 起别名的类型 也是可以有方法的
func (c A) test() {
	fmt.Println(c)
}

func (m Maps) test1() {
	fmt.Println(m)
}

func (s Student) test() {
	s.Name = "101111"
	fmt.Println(s.Name)
}

func (s *Student) test2() {
	s.Name = "10"
	fmt.Println(s.Name)
}
