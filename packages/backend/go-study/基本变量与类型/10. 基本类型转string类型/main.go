package main

import (
	"fmt"
	"strconv"
)

func main() {

	// fmt.Sprintf("%v", a) 进行转换
	var a int8 = 99 // 其余类型同样的  bool float

	s := fmt.Sprintf("%v", a)

	fmt.Printf("%T %v", s, s) // string 99

	// strconv 进行转换 不常用

	s1 := strconv.FormatInt(int64(a), 10) // 这个函数必须传入一个int64，和一个进制数

	fmt.Printf("%T %v", s1, s1)

}
