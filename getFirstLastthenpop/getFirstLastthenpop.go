package main

import "fmt"

/* 4,5 */
func main() {
	arr := []int{1, 2, 3, 4, 5, -100}
	fmt.Println(arr)
	slast := arr[:]
	fmt.Println(slast[len(slast)-1])
	slast = slast[0 : len(slast)-1]
	fmt.Println(slast)
	fmt.Println("-------------------")
	arr = []int{999, 2, 3, 4, 5, 6}
	slast = arr[:]
	fmt.Println(slast)
	fmt.Println(slast[0])
	copy(slast[0:], slast[1:])
	slast = slast[0 : len(slast)-1]
	fmt.Println(slast)
}
