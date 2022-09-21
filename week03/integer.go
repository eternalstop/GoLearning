package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x1 int
	var x2 int32
	var x3 int64
	var x4 uint32
	var x5 uint64
	var y1 int8
	var y2 uint8

	y1 = 127
	y1 = y1 + 1
	y2 = 255
	y2 = y2 + 1

	fmt.Println(x1, unsafe.Sizeof(x1))
	fmt.Println(x2, unsafe.Sizeof(x2))
	fmt.Println(x3, unsafe.Sizeof(x3))
	fmt.Println(x4, unsafe.Sizeof(x4))
	fmt.Println(x5, unsafe.Sizeof(x5))
	fmt.Println(y1, y2)
}
