package main

import "fmt"

func main() {

	// -------------------------------
	// 值传递和引用传递
	var a int = 10
	var b int = 20
	fmt.Println(a, b)
	exchageNum(a, b)
	fmt.Println(a, b) // 没变化

	exchageNum1(&a, &b)
	fmt.Println(a, b) // 没变化

	// ------------------------------------
	// 调用可变参数
	test(1, 2, 3, 4, 5)

	// --------------------------------------

	// 函数赋给一个变量
	f := test
	f(1, 2, 3, 4, 5, 6)

	// --------------------------------------

	// 自定义数据类型
	type MyFunc func(args ...int) int

	var n MyFunc = test
	n(1, 2, 3, 4, 5, 6)

	type Myint int // 其实就是起了一个别名

	var n1 Myint = 100
	var n2 int = 200

	//不能将 n2 赋值个n1 类型不相同 会报错
	// n1 = n2

	// 但是可以强制转换
	n1 = Myint(n2)
	n2 = int(n1)

	// ------------------------------------

	// 函数作为参数传入到另一个函数中进行调用
	test1(100, test2)

}

// 可以被其他包调用 public 类型
func AddNum(a, b int) int {
	return a + b
}

// 不能被其他包调用  private 类型
func addSum(a, b int) int {
	return a + b
}

// 值传递
func exchageNum(a, b int) {
	var c int

	c = a
	a = b
	b = c
}

// 引用传递
func exchageNum1(a *int, b *int) {
	var c int

	c = *a
	*a = *b
	*b = c
}

// 可变参数函数

func test(args ...int) int {
	fmt.Println(len(args))
	fmt.Println(cap(args))
	for i, v := range args {
		fmt.Println(i, v)
	}

	return 0
}

func test2(c int) {
	fmt.Println(c)
}
func test1(a int, b func(c int)) {
	b(a)
}
