package main

import "fmt"

func intSeq() func() int {
	fmt.Println("This is an outer function..............!")
	i := 0
	fmt.Println(i)
	return func() int {
		fmt.Println("This is an inner function.............!")
		i++
		fmt.Println(i)
		return i
	}
}

func main() {
	fmt.Println("This is main function...............!")
	nextInt := intSeq()
	fmt.Println(nextInt())
}
