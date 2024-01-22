### 简介

- fmt.Sprintf()函数用于格式化输出字符串，格式化字符串的语法和C语言的printf函数一致，但是Go语言的格式化字符串语法比C语言的语法要简单很多。
  - fmt.Sprintf()函数返回一个字符串，该字符串是格式化后的字符串。 可以用来将其余类型转为字符串
  - `a := fmt.Sprintf("Hello %v", 32)`
  - `a := fmt.Sprintf("Hello %s", "world")`
  - `b := fmt.Sprintf("Hello %T %v", a, a)` 
  
- strconv 包内的函数进行转换, 不如上面的方式方便，但是如果需要转换进制之类的就比较方便
  - `s1 := strconv.FormatInt(int64(a), 10)` int 转 string  这个函数必须传入一个int64，和一个进制数
