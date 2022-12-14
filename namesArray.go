package main

import (
	"fmt"
)

var namesArray [5]string

func main() {
	for i := 0; i < 5; i++ {
		val := ""
		fmt.Println("i: ", i)
		fmt.Scan(&val)
		namesArray[i] = val
	}
	fmt.Println(namesArray)
}
