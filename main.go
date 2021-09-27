package main

import (
	"fmt"
)

func Add(x int, y int) (result int) {
	return x + y
}

func main() {
	fmt.Println(Add(2, 3))
}
