package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// func main() {
func echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	var sep = flag.String("s", " ", "separator")
	var line = flag.Bool("n", false, "newline")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if *line {
		fmt.Println()
	}
	fmt.Print("L")
}
