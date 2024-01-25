package main

import "fmt"

func main() {

	// 定义一个数组
	var arr [3]int // 这个数组存储三个int类型的值, int类型没有初始化 默认值就为 0  [0,0,0]

	fmt.Println(arr)

	ar1 := test(arr) // 值传递
	fmt.Println(ar1)
	fmt.Println("arr = ", arr)

	// 获取数组长度
	fmt.Println(len(arr))

	// 可以使用 ... 来表示 省略数组长度，通过初始化的数据 自动推断长度
	var arr1 = [...]int{1, 2, 3, 2, 5} // {xx,xx} 数组初始化
	fmt.Println(len(arr1))             // 5

	var arr2 = [...]string{} // string 类型的数组，默认值都为空
	fmt.Println(len(arr2))

	var arr3 = [...]float32{3.3, 4.4, 5.5}
	fmt.Println(len(arr3))

	// 数组初始化，用下标
	var arr10 = [...]int{0: 1, 3: 10} // 没有设置的就是默认值0
	fmt.Println(arr10)

	// 普通for循环，和 for range 对数组
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for _, v := range arr2 {
		fmt.Println(v)
	}

	// 多维数组
	var arr4 = [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	// 多重循环数组
	for _, v := range arr4 {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
	test2(&arr10)
}
func test(ar [3]int) [3]int {

	ar[0] = 1 // 这里修改不影响 原数组arr

	return ar
}

func test2(ar *[4]int) {
	(*ar)[0] = 1 // 指针类型将，原数组被修改
}
