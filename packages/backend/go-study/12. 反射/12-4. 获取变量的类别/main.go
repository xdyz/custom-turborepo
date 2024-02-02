package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	stu := Student{
		Name: "张三",
		Age:  18,
	}

	useReflect(stu)

}

// 对哦结构体 进行反射
func useReflect(v interface{}) {
	fmt.Println(reflect.TypeOf(v))                                // [3]int
	fmt.Printf("reflect.TypeOf(v) 返回类型 %T \n", reflect.TypeOf(v)) // 输出 *reflect.rtype
	fmt.Println(reflect.ValueOf(v))                               //
	fmt.Printf("reflect.ValueOf(v) 返回类型 %T \n", reflect.ValueOf(v))

	// 转回来

	vIn := reflect.ValueOf(v).Interface() // 将 reflect.ValueOf(v) 的 reflect.Value 类型 转为 interface{} 类型

	n, flag := vIn.(Student) // 对interface{} 类型断言是不是一个Student
	if flag {
		fmt.Println(n)
	} else {
		fmt.Println("不是Student类型")
	}

	fmt.Println(reflect.ValueOf(v).Kind()) // 返回一个 struct 类别   Student 属于自定义类型
}
