package week03

import (
	"fmt"
)

func study_string() {
	var s1 string
	var c string
	var b byte
	s1 = "hello world"
	b = s1[0]
	c = s1[0:3]
	fmt.Println(s1)
	fmt.Println(b)
	fmt.Println(c)
}
