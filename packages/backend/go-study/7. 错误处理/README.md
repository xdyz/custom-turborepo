### 简介



#### panic
执行出错，后面的代码不再执行，直接返回，不会继续执行





### 错误处理/捕获机制

go 中没有 try catch 来捕获错误 而是用其余的方式

#### defer 和 recover 来捕获错误

defer 会在函数执行结束时执行，无论函数执行成功还是失败

如果recover 是在defer 函数外被调用，recover 会捕获 panic，如果函数执行出错，会返回 recover 的值，不会继续执行函数后面的内容，如果此时没有panic，recover 会返回nil




### 自定义错误，自己抛出错误

