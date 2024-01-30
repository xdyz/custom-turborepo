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

	fmt.Println(m5)

}
