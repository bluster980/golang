package main

import (
	"fmt"
)

// Here interface contains the nethods signatures basically the top view.
type myType interface {
	multiply() int
}

// here struct cover the composite data type as a multiple variable.
type myStruct struct {
	x int
	y int
}

// method to do subtraction
func (m myStruct) subtract() int {
	return m.x - m.y
}

// method to do multiplication
func (m myStruct) multiply() int {
	return m.x * m.y
}

// a function to pass arguments to multiply method
func multiplication(item myType) int {
	return item.multiply()
}

// a function to addition of two numbers
func addition(x int, y int) (int, error) {
	return x + y, nil
}

func main() {
	fmt.Println("Hello, World!")
	dummy, err := addition(5, 6)
	fmt.Println(dummy, err)
	temp := myStruct{5, 6}
	fmt.Println(temp.subtract())
	dumb := myStruct{5, 6}
	fmt.Println(multiplication(dumb))
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
