package main

import "fmt"

func Zero(x int) {
	x = 0
}

func main() {
	x := 5
	Zero(x)
	fmt.Println(x)
}
