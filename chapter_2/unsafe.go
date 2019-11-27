package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value int64 = 5
	var p1 = &value
	// This line allows for p2 to reference an int64 pointer as int32
	var p2 = (*int32)(unsafe.Pointer(p1))

	fmt.Println("*p1: ", *p1)
	fmt.Println("*p2: ", *p2)

	*p1 = 543412341231231212
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)

	*p1 = 54341234
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
}
