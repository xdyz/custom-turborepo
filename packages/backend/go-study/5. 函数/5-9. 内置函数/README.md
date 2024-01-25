### 简介

go 内置的函数 ，不需要导入包


#### len 函数
返回变量的字节数，中文占3个字节


#### new 函数
new 函数返回一个指向类型为 T 的零值的指针，T 必须是引用类型。

一般 值类型 int float bool string 数组 结构体 都用new 来分配内存

引用类型 map slice channel interface 用 make 函数来分配内存