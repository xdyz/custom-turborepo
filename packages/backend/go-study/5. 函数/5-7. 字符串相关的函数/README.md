### 简介

字符串函数属于系统函数，如果不会 就去官网查就好了 没必要全部死记硬背
其余系统函数同理

#### len(str)
统计字符串的长度，按字节统计

#### r := []rune(str)
字符串遍历

#### n, err := strconv.Atoi(str)
字符串转整数

#### str = strconv.Itoa(n)
整数转字符串

#### strings.Contains(str, substr)
判断字符串是否包含子串

#### strings.EqualFold(str1, str2)
判断两个字符串是否相等，忽略大小写

#### strings.Index(str, substr)
返回子串在字符串中第一次出现的位置, 如果没有就返回 -1


#### strings.Replace(str, old, new, n)  
将 str 中 old 替换为 new, n 指定替换的次数
n 如果为 -1, 则替换所有出现的 old
n 为 0 就是不替换
n 为 几 就是替换几个

#### strings.Split(str, sep)
将字符串 str 按照 sep 分割成一个字符串数组

#### strings.ToLower(str)  strings.ToUpper(str)
大小写转换


#### strings.TrimSpace(str)
删除字符串 str 中的空格 左右两边的空格


#### strings.Trim(str, cutset)
删除字符串 str 中的字符 cutset 左右两侧，最左边和最右边


#### strings.HasPrefix(str, prefix)
判断字符串 str 是否以 prefix 变量内容 开头

#### strings.HasSuffix(str, suffix)
判断字符串 str 是否以 suffix 变量内容 结尾