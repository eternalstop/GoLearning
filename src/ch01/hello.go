package main

import (
	"fmt"
)

func main() {
	var x int
	x = 1
	var y int
	y = 2
	swap(&x, &y)
	fmt.Println(x, y)
	//fmt.Println("hello, golang world!!!", os.Args[1])
}

func swap(p *int, q *int) {
	var t = *p
	*p = *q
	*q = t
}
