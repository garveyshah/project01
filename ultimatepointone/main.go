package main

import (
	"fmt"
	"project01"
)

func main() {
	a := 0
	b := &a
	n := &b
	project01.UltimatePointOne(&n)
	fmt.Println(a)
}
