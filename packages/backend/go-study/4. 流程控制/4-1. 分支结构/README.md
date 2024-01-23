### 简介

主要分为下面两类

- if 分支
  - 单分支
    - 条件语句 的() 可以不写， {} 必须写
  ```golang
  if 条件语句 {
     逻辑语句
  }
  ```
  - 多分支
  - 双分支

- switch case
  - go 中的 case 不需要break， 默认会 break， 会自动跳出
  - fallthrough 会穿透当前case 进入到下一个case 如果下一个case 没有 fallthrough 会跳出 如果有 就继续穿透， 相当于JS 中的 switch 不带 break
  - case 可以使用表达式 或者变量，或者有返回值的函数，单个case 多个值时， 用逗号隔开
  - switch 后面可以不跟上变量，函数带返回值的，