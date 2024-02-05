package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name    string
	Age     int
	address string
}

func (s Student) WaitGroupetName() string {
	return s.Name
}
func (s Student) Print() {
	fmt.Printf("调用了Print方法\n")
	fmt.Println("name: ", s.Name)
}

// 这个无法被 NumMethod 方法获取到
func (s Student) getAddress() string {
	return s.address
}

func (s Student) Set(name string, age int) {
	s.Name = name
	s.Age = age

	fmt.Println("修改后的值为：", s)
}

func main() {
	std := Student{
		Name:    "张三",
		Age:     18,
		address: "北京",
	}

	useReflect(std)
	fmt.Println("修改后的值为：", std)

	// 如果要修改值，就必须传入指针
	useReflectPoint(&std)
	fmt.Println("修改后的值为：", std)
}

// 不修改 原 结构体的值，因为使用的时候 传入的是一个结构体，而不是指针
func useReflect(v interface{}) {
	rv := reflect.ValueOf(v) // 先转为 reflect.Value 类型

	// 获取 v 的属性个数
	num := rv.NumField() // 返回个int
	fmt.Println("num:", num)

	// 所有属性，都存在一个数组中，可以通过下标访问到
	for i := 0; i < num; i++ {
		fmt.Printf("第 %d 个参数的值为 %v \n", i, rv.Field(i)) // Field(i) 返回参数的值
	}

	// 获取方法的个数

	num2 := rv.NumMethod()     // 这个只能获取到首字母为大写的方法的个数
	fmt.Println("num2:", num2) // 3 而不是4

	// 方法同上，只不过方法的 下标顺序为 方法首字母的asscii码值

	// 准确的调用方法
	//
	rv.MethodByName("Print").Call([]reflect.Value{}) // Call 接收一个参数，参数为 []reflect.Value 切片 类型

	// 先定义一个 reflect.Value 的切片类型

	var args []reflect.Value = []reflect.Value{
		reflect.ValueOf("李四"), // 将值转换为 reflect.Value 类型
		reflect.ValueOf(19),   // 同上
	}

	rv.MethodByName("Set").Call(args)
}

func useReflectPoint(v interface{}) {
	// 先转为 reflect.Value 类型
	rv := reflect.ValueOf(v) // 先转为 reflect.Value 类型

	rv.Elem().FieldByName("Name").Set(reflect.ValueOf("王五")) // Elem() 返回的是指针的值，所以要调用 Elem() 方法获取指针的值
	rv.Elem().FieldByName("Age").Set(reflect.ValueOf(20))
	// rv.Elem().FieldByName("address").Set(reflect.ValueOf("上海")) // 这个不能修改，因为是小写字母开头
}
