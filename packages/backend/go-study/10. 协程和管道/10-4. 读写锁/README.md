### 简介
读写锁 RWMutex 是一个 适用于 读比较多，写比较少的场景


### 读写锁的使用
```go
var rwMutex sync.RWMutex

func read() {
  rwMutex.RLock() // 读加锁
  rwMutex.RUnlock() // 读解锁
  
  rwMutex.WLock() // 写加锁
  rwMutex.WUnlock() // 写解锁
}