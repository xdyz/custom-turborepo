package main

import "fmt"

func main() {
	var flag bool = true

	var flag1 bool = false

	flag3 := false

	if flag {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	switch flag1 {
	case true:
		fmt.Println("true")
		break
	case false:
		fmt.Println("false")
		break
	default:
		fmt.Println("default")
		break
	}

	switch flag3 {
	case flag1:
		fmt.Println("true")
		break
	case flag:
		fmt.Println("false")
		break
	default:
		fmt.Println("default")
		break
	}
}
