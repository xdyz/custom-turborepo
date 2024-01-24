// package main

// import "fmt"

// func main() {
// 	var a int = 1
// 	var b int = 2
// 	fmt.Println(add(&a, &b)) // 最后执行
// 	fmt.Println(a, b)
// }

// func add(a *int, b *int) int {

// 	defer fmt.Println("*a = ", *a) // 先进入栈，后被执行弹出
// 	defer fmt.Println("*b = ", *b) // 后进入栈，先被执行弹出

// 	// 上难度， 此时 如果给 a  和 b += 值 会怎么变化？？

// 	// 此时 上面两个defer 的输出不会改变的

// 	*a += 100
// 	*b += 48

// 	fmt.Println("a + b = ", *a+*b) // 最先执行

// 	return *a + *b // defer 在return 之后执行
// }

// // 上述的执行结果为
// // a + b = 3
// // b = 2
// // a = 1
// // 3

// 下面的这个例子 叫做 命名返回值 函数 就是 返回的值 已经被命名了，可以在defer中修改
// 然后函数的 return 什么东西 根本就影响不到了

package main

import "fmt"

// 这里可以看出 defer 和 return 的执行顺序 且 defer 可以改变输出的值
// 好像和上面的例子 说的不太一样

// result 是一个在返回前 就已经被明明的变量了，相当于在函数外层定义的普通变量，所以defer这里能够改变其值
// 这里的result 返回值 相当于在外面定义好了一个变量了
func foo() (result int) {
	defer func() {
		result = result * 2
		fmt.Println("Inside defer, result is modified to:", result)
	}()

	// 这里先给变量result 赋值 5
	result = 5
	fmt.Println("Before return, result is:", result)

	// return 此时就可以看成是两个步骤
	// 1. 这里返回值赋给result
	// 2. 等前面都执行完  defer 也执行完了， 在执行第二步操作，返回

	// result 相当于是 外部定义的变量，此时第二步执行前 就已经被defer 给改变了
	// 所以 return 后面的返回 根本就无意义了

	// 此时这里无轮返回啥都不影响 外部finalResult
	return
}

func main() {
	finalResult := foo()
	fmt.Println("After return, finalResult is:", finalResult)
}
