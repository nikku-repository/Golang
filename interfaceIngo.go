package main

import (
	"fmt"
	"reflect"
)

type Car interface {
	details()
}

func show_data(c Car) {
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))
}

type alto struct {
	modelNumber int
	regNo       float64
}

func (a alto) details() {
	fmt.Println("This is an alto car.......")
	fmt.Println(a.modelNumber, a.regNo)
}

type fortuner struct {
	modelNumber int
	regNo       float64
}

func (f fortuner) details() {
	fmt.Println("This is a fortuner car.......")
	fmt.Println(f.modelNumber, f.regNo)
}

func main() {
	a := alto{12345, 6.4423}
	a.details()
	f := fortuner{1025, 3.1425}
	f.details()
	show_data(a)
	show_data(f)
}
