package main

import "fmt"

func main() {
	// 创建一个 key为 int, value 为string 的map
	// 此时只声明map,但是是不分配内存的
	var m1 map[int]string // 此时 m1 == nil

	// 如果想要在声明的时候分配内存，需要使用make 来创建
	// size 可以不传 默认分配 1个空间
	m1 = make(map[int]string, 10) // 10 代表这个map 可以存储10个键值对, 会自动扩容

	m1[1] = "a"
	m1[2] = "b"
	m1[3] = "c"
	m1[10] = "c"
	m1[5] = "c"
	m1[9] = "c"
	m1[8] = "c"
	m1[7] = "c"
	m1[6] = "c"
	m1[4] = "c"
	m1[8] = "c"
	m1[11] = "c"

	fmt.Println(m1)

	// 顺序都是无序的，每次都不一样
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	m3 := map[string]string{
		"name": "fa",
		"age":  "18",
		"sex":  "man",
	}

	fmt.Println(m3)

	// 删除 m3 中 键值对
	delete(m3, "sex")

	fmt.Println(m3)

	m4, hasKey := m3["sex"]

	fmt.Println(m4, hasKey)

	fmt.Println(len(m4))

	// map 嵌套map
	m5 := map[string]map[string]string{
		"fa": {
			"name": "fa",
			"age":  "18",
			"sex":  "man",
		},
		"fe": {
			"name": "fe",
			"age":  "18",
			"sex":  "man",
		},
	}

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
