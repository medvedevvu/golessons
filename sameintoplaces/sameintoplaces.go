// задание 3
package main

import "fmt"

func main() {
	firstSlice := []int{1, 3, 11, 5, 6, 6, 7, 8, 9}
	secondSlice := []int{0, 3, 11, 5, 6, 12, 7, 8, 9}
	resmap := make(map[int]int)
	firstSlice = append(firstSlice, secondSlice[:]...)
	for _, value := range firstSlice {
		resmap[value]++
	}

	for idx, value := range resmap {
		if value > 1 {
			fmt.Println(idx)
		}
	}
}
