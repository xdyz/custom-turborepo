package main

import (
	"fmt"
)

func main() {
	// for 循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for 循环 初始化值可以写到外面
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// for 可以 作为 while循环使用
	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	// for 无条件 一直执行下去
	// for {
	// 	fmt.Println("无条件")
	// }

	// for  range 循环一个数组
	for i, v := range []int{1, 2, 3, 4, 5} {
		fmt.Println(i, v)
	}

	// for range 循环一个切片
	for i, v := range []string{"a", "b", "c", "d", "e"} {
		fmt.Println(i, v)
	}

	// for  range 循环一个字符串 此时 v 是一个字符，需要用 %c 才显示具体字符
	for i, v := range "abcde" {
		// fmt.Println(i, v)
		// fmt.Printf("%v %q \n", i, v)
		fmt.Printf("%v %c \n", i, v)
	}

	// for range 循环一个 中文字符串
	// 中文占两个字节 所以 i  是 0, 3
	for i, v := range "你好" {
		// fmt.Println(i, v)
		fmt.Printf("%v %c \n", i, v)
	}

	// for range 循环一个 map
	for k, v := range map[string]string{"a": "1", "b": "2", "c": "3"} {
		fmt.Println(k, v)
	}

	// 用break 退出循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 只退出循环
		}
		fmt.Println(i)
	}

	// 双 for循环 break + label 退出
label1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			if i == 5 {
				break label1 // 退出双 for循环 到代码块label1： 就变成无线循环了
			}
		}
		fmt.Println(i)
	}

	// 使用 continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	// 写一个双重for循环 continue + label
label3:
	for i := 0; i < 10; i++ {
		for j := 0; j < 30; j++ {
			if i == 5 {
				continue label3 // 退出双 for循环 到代码块label3： 就变成无线循环了
			}
		}
		fmt.Println(i)
	}

	// 使用 goto
label2:
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto label2 // 到指定代码块位置
		}
		fmt.Println(i)

		// end:
		// fmt.Println("end")
	}
	// end:
	// 	fmt.Println("end")

	// for 和 return
	for i := 0; i < 10; i++ {
		if i == 5 {
			return // 直接跳出函数，不止跳出循环
		}
		fmt.Println(i)
	}

}
