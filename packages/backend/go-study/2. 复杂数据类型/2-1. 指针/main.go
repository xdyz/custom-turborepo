package main

import "fmt"

func main() {
	var a int = 10

	fmt.Println(&a) // a的值 所存储在的地址

	var p *int = &a

	// a 的值 可以用两种方式表示
	fmt.Println("用 a  表示 ", a)
	fmt.Println("用 *p 表示 ", *p)

	// a 的地址 可以用两种方式表示
	fmt.Println("用 &a 表示 ", &a)
	fmt.Println("用 p  表示 ", p)

	// p 本身也是开辟了一个新内存地址，其中存储着 a 的地址 我们还可以查看 p 的地址

	fmt.Println("用 &p 表示 ", &p) // 与上面的值不同

	// *p 值的改变 也会导致 a 的值改变
	*p = 20
	fmt.Println("用 a  表示 ", a) // 变成了20
	fmt.Println("用 *p 表示 ", *p)
}
