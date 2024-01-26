package test

import "fmt"

type Person struct {
	Age int
}

type Student struct {
	Age int
}

func Test() {

	var s Student
	var p Person

	// s = p // 这样不能直接转换，因为类型不一致，需要进行类型转换

	s = Student(p) // 转换必须保证 属性的 类型，名称，个数相同才可以

	fmt.Println(s)

	// 利用type 定义一个 结构体的别名
	// type Student2 Student 这种别名下，无法直接转换，必须强转
	type Student2 = Student

	var s2 Student2
	s2 = s

	fmt.Println(s2)
}
