// --- задание 4
package main

import "fmt"

func fiba(idx int) int {
	fimap := map[int]int{0: 1, 1: 1}

	switch idx {
	case 0:
		fallthrough
	case 1:
		return fimap[idx]
	default:
		for i := 2; i <= idx; i++ {
			fimap[i] = fimap[i-1] + fimap[i-2]
		}
		return fimap[idx]
	}

}

func main() {
	for idx := 0; idx < 10; idx++ {
		fmt.Println(fiba(idx))
	}
}
