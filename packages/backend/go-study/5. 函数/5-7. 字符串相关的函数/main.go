package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// go 语言中 utf-8 字节符号，所以一个中文汉字是 3 个字节
	var a string = "golang你好a"
	fmt.Println(len(a)) // 12

	// 字符串遍历 for range 方法
	for i, v := range a {
		fmt.Printf("%d: %c\n", i, v)
	}

	// 0: g
	// 1: o
	// 2: l
	// 3: a
	// 4: n
	// 5: g
	// 6: 你
	// 9: 好
	// 12: a

	// 第二种字符串的遍历方法 利用切片类型
	r := []rune(a)
	for i, v := range r {
		fmt.Printf("%d: %c\n", i, v)
	}

	// 0: g
	// 1: o
	// 2: l
	// 3: a
	// 4: n
	// 5: g
	// 6: 你
	// 7: 好
	// 8: a

	// 字符串转整数

	var s string = "123123"
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	// 整数转字符串
	fmt.Println(strconv.Itoa(124443123))

	// 判断字符是否包含子串
	fmt.Println(strings.Contains(s, "123")) // 返回一个bool 值

	// 统计字符串中包含几个指定的子串
	fmt.Println(strings.Count(s, "123")) // 返回一个int 值

	// 比较两个字符串，不区分大小写 是否相等
	fmt.Println(strings.EqualFold("abc", "ABC")) // 返回一个bool

	// 返回子串在字符串中第一次出现的索引位置
	fmt.Println(strings.Index("124443123", "123")) // 6 找不到就返回-1

	// 字符串的替换
	fmt.Println(strings.Replace("124443123", "123", "asa", 1))

	// 字符串切割
	fmt.Println(strings.Split("124-443-123", "-"))
	fmt.Println(strings.Split("124443123", "1"))

	// 字符串大小写转换
	fmt.Println(strings.ToLower("abc"))
	fmt.Println(strings.ToUpper("abc"))

	// 删除字符串的空格（左右两侧的，中间的不管）
	fmt.Println(strings.TrimSpace(" 12 3  "))

	// 删除两侧指定字符
	fmt.Println(strings.Trim("~12 3  ", "~"))

	// 判断字符串是否以xxx开头
	fmt.Println(strings.HasPrefix("124443123", "123"))

	// 判断字符串是否以xxx结尾
	fmt.Println(strings.HasSuffix("124443123", "123"))

}
