// задания 1
package main

import (
	"fmt"
)


func main() {
	mslice := make([]int, 10)
	/* инит slice */
	for idx := 0; idx < len(mslice); idx = idx + 1 {
		mslice[idx] = 9
	}
	fmt.Print(mslice)
	/* прибавляю 1 к каждому элементу */
	for idx := 0; idx < len(mslice); idx = idx + 1 {
		mslice[idx] += 1
	}
	fmt.Print(mslice)
}
