### 简介

map 映射 是go 中一种内置的类型，通过键值对的方式存储数据，


#### 基本语法

- key 一般使用 int string 类型, 不可以为 slice map 数组
- value 通常为 数字，string, map， 结构体， bool

#### 注意
- 只声明的map是没有分配内存空间的
- key 不可以重复，重复的key 赋值value 时 ，后面会覆盖前面的
  - value 可重复
- map 的 key value 是无序的，所以map 遍历时 并不会一定按照 key 的顺序遍历
  - 所以每次遍历的顺序都是不同的
- make 函数的第二个参数可以省略，默认为 1 个空间大小


#### 初始化
```go
m3 := map[string]string{
		"name": "fa",
		"age":  "18",
		"sex":  "man",
}
```

### map 的 操作

#### 增加/更新
- map[key] = value
  - 如果 key 不存在，会自动创建 key 然后value赋值
  - 如果 key 存在，会更新 value

#### 删除
需要使用内置函数 delete 接收两个参数，第一个参数为 map 对象，第二个参数为 key
- 如果key存在，会删除键值对
- 如果key不存在，也不会报错


#### 清空
- go中没有提供单独清空的操作，需要遍历然后对键值对逐个删除 一般不使用
- 给变量 重新去make 一个新的map，旧的就会被GC 自动回收


#### 查找
value, bool = map[key] 
value 为返回的值，bool 为 true 表示存在，false 表示不存在


#### 获取长度

len(map)