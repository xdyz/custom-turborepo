### 简介

#### 一次性读取

读取文件的内容并显示在终端，一次性将文件的内容全部读取到内存中，然后输出到终端。
这种对方式，适用于文件不大的情况，如果文件很大，会撑爆内存，所以不推荐使用。

现在不适用 ioutil.ReadFile 读取文件，而是使用 os.ReadFile 读取文件



#### 带缓冲的读取
通过 bufio.NewReader 读取文件，可以指定缓冲区大小，从而控制读取文件的大小。 
默认的缓冲区大小为 4096，即 4KB。
这个方法，没有封装文件的开关，所以需要自己控制文件的打开和关闭