package main

import (
	"fmt"
)

func getItem(i int, pslc []int) {
	if i >= cap(pslc) || i < 0 {
		return
	}
	fmt.Println(pslc)
	fmt.Println(pslc[i])
	ln := len(pslc) - 1
	tmp := make([]int, ln)
	copy(tmp, pslc[0:i])
	copy(tmp[i:], pslc[i+1:])
	fmt.Println(tmp)
}

/* 6 */
func main() {
	slc := ([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})[:]
	getItem(1, slc)
	fmt.Print("Ok")
}
