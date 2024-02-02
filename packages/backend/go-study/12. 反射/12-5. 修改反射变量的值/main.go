package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 对 int 类型的变量进行反射
	var a int = 3
	useReflect(&a)
	fmt.Println(a)

	// // 对 string 类型的变量进行反射
	// var b string = "hello"
	// useReflect(&b)

	// // 对 float 类型的变量进行反射
	// var c = 3.14
	// useReflect(&c)

	// // 对 bool 类型的变量进行反射
	// var d = true
	// useReflect(&d)
	// fmt.Println(d)

	// // 对 [n]type 数组类型的变量进行反射
	// var e = [3]int{1, 2, 3}
	// useReflect(&e)
	// fmt.Println(e)

}

func useReflect(v interface{}) {
	// fmt.Println(reflect.TypeOf(v))                                // [3]int
	// fmt.Printf("reflect.TypeOf(v) 返回类型 %T \n", reflect.TypeOf(v)) // 输出 *reflect.rtype
	// fmt.Println(reflect.ValueOf(v))                               //
	// fmt.Printf("reflect.ValueOf(v) 返回类型 %T \n", reflect.ValueOf(v))

	rv := reflect.ValueOf(v) // 返回一个 reflect.Value 类型
	t := fmt.Sprintln(reflect.TypeOf(v))
	fmt.Println(t)

	rv.Elem().SetInt(3333333)
}
