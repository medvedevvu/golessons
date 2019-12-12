package main

import "fmt"

func seven() {
	sl1 := []int{1, 2, 3, 4}
	sl2 := []int{5, 6, 7, 8}
	sl1 = append(sl1, sl2...)
	fmt.Println(sl1)
}

func eight() {
	sl1 := []int{1, 6, 7, 6}
	sl2 := []int{5, 6, 7, 8}
	sl3 := []int{}
	fmt.Print(sl1)
}

func main() {
	seven()
	eight()
}
