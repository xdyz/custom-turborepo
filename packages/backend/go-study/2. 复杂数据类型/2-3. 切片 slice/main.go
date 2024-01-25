package main

import "fmt"

func main() {

	// 从数组中切出一个切片
	var a = [5]int{1, 2, 3, 4, 5}

	var b = a[1:3] // 下标1 开始 到下标3 不包括小标3 的一段数据 [2,3]

	// 4 就是设置的新切片的容量，但是超过了长度，此时是不能超过的，所以容量和长度相等都为3
	// 4 位置的这个值不可以小于 长度，不然会报错，所以这里的情况是不能小于3
	c := a[1:3:3] // 第三个数 是这个切片的容量
	fmt.Printf("%T %v", b, b)
	fmt.Printf("%T %v %v", c, c, cap(c))

	// make 切片

	d := make([]int, 3, 5)

	// 直接定义切片

	var e = []int{1, 2, 3}

	// nil 空切片
	var f []int

	fmt.Printf("%T %v", d, d)
	fmt.Printf("%T %v", e, e)
	fmt.Printf("%T %v", f, f)

	// 切片的拷贝， 互相之间不受影响
	var g = []int{1, 2, 3}
	h := make([]int, len(g))
	copy(h, g) // 将g的值复制到h中， 深拷贝

}
