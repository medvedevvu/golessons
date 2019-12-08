package main

import (
	"fmt"
	"strconv"
)

func swap(x int, y int) (int, int) {
	return y, x
}

func main() {
	x, y := 2, 5
	fmt.Printf("x=%s y=%s \n", strconv.Itoa(x), strconv.Itoa(y))
	x, y = swap(x, y)
	fmt.Printf("x=%s y=%s \n", strconv.Itoa(x), strconv.Itoa(y))
}
