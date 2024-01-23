### 简介


有两种循环

- for 循环
  - 循环变量可以写到外面
  - 初始化值 不可以在 for 后 用var 定义
  - 省略迭代，可以像while循环一样
  - for 循环可以嵌套
  - for 循环没有任何条件，将作为永真循环
- for range 循环
  - 可以用于遍历数组，切片，字符串，map, channel
  - 遍历数组时，可以同时获取索引和值
  - 遍历切片时，同上
  - 遍历字符串时，同上
  - 遍历map时，同时获取key和value
  - 遍历channel时，同时获取value和done


- 关键字
  - break   退出循环， 停止是正在执行的离break 最近的循环，如果是多循环嵌套的话
  - break 还可以加 label ，跳到指定的循环
  - continue 跳过本次循环, 不执行continue 后面的语句，直接进入下次循环
  - continue 还可以加 label ，跳到指定的循环
  - goto 跳转到指定的代码块, 一般不建议使用
  - return  直接跳出函数，不止跳出循环