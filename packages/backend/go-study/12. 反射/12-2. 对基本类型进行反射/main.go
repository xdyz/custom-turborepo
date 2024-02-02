package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 对 int 类型的变量进行反射
	var a int = 3
	useReflect(a)
	// fmt.Println(reflect.TypeOf(a))          // 输出 int
	// fmt.Printf("%T \n", reflect.TypeOf(a))  // 输出 *reflect.rtype
	// fmt.Println(reflect.ValueOf(a))         // 输出 3 只是输出3 但是不能直接当做数字使用
	// fmt.Printf("%T \n", reflect.ValueOf(a)) // 输出 reflect.Value    reflect.ValueOf(a) 返回的 类型是  reflect.Value

	// 对 string 类型的变量进行反射
	var b string = "hello"
	useReflect(b)
	// fmt.Println(reflect.TypeOf(b))
	// fmt.Printf("%T \n", reflect.TypeOf(b)) // 输出 *reflect.rtype
	// fmt.Println(reflect.ValueOf(b))
	// fmt.Printf("%T \n", reflect.ValueOf(b))

	// 对 float 类型的变量进行反射
	var c = 3.14
	useReflect(c)
	// fmt.Println(reflect.TypeOf(c))
	// fmt.Printf("%T \n", reflect.TypeOf(c)) // 输出 *reflect.rtype
	// fmt.Println(reflect.ValueOf(c))
	// fmt.Printf("%T \n", reflect.ValueOf(c))

	// 对 bool 类型的变量进行反射
	var d = true
	useReflect(d)
	// fmt.Println(reflect.TypeOf(d))
	// fmt.Printf("%T \n", reflect.TypeOf(d)) // 输出 *reflect.rtype
	// fmt.Println(reflect.ValueOf(d))
	// fmt.Printf("%T \n", reflect.ValueOf(d))

	// 对 [n]type 数组类型的变量进行反射
	var e = [3]int{1, 2, 3}
	useReflect(e)
	// fmt.Println(reflect.TypeOf(e))         // [3]int
	// fmt.Printf("%T \n", reflect.TypeOf(e)) // 输出 *reflect.rtype
	// fmt.Println(reflect.ValueOf(e))        //
	// fmt.Printf("%T \n", reflect.ValueOf(e))
}

func useReflect(v interface{}) {
	fmt.Println(reflect.TypeOf(v))                                // [3]int
	fmt.Printf("reflect.TypeOf(v) 返回类型 %T \n", reflect.TypeOf(v)) // 输出 *reflect.rtype
	fmt.Println(reflect.ValueOf(v))                               //
	fmt.Printf("reflect.ValueOf(v) 返回类型 %T \n", reflect.ValueOf(v))
}
