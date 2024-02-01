### 简介
1. select 是Go中的一个控制语句， 类似于switch语句，用于处理异步IO操作，select会监听case 语句中channel的读写操作，当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的操作
   - select 中的语句必须是一个 channel 通道操作
   - select 中的default 子句总是可运行的
2. 如果有多个case都可以运行，select会随机公平的选出一个执行，其他的不会执行
3. 如果没有可运行的case语句，且有default语句，那么就会执行default的动作
4. 如果没有可运行的case语句，且没有default语句，那么select 将会阻塞，直到某个case通信可以运行