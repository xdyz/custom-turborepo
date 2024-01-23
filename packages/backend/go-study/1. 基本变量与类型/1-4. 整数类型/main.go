package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Print(unsafe.Sizeof(65535))
}
