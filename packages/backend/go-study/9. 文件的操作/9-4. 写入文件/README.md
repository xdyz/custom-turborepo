### 简介

程序将数据流写入到 一个数据源中


#### 打开文件
`os.OpenFile(name string, flags int, perm os.FileMode) (*File, error)`


- name: 文件名/地址
- flags: 如果有多个一起使用 需要用 | 连接即可
  - os.O_RDONLY: 只读
  - os.O_WRONLY: 只写
  - os.O_RDWR: 读写
  - os.O_APPEND: 以追加的方式打开文件
  - os.O_CREATE: 如果文件不存在则创建
  - os.O_EXCL: 如果文件已存在则报错
  - os.O_SYNC: 同步IO
  - os.O_TRUNC: 如果文件已存在则清空内容
- perm : 文件权限 windows 下 写不写都无效