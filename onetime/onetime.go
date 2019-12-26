// задание 2

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	max := 99
	min := 11
	bslice := [...]int{99: 0}
	for idx, _ := range bslice {
		bslice[idx] = rand.Intn(max-min+1) + min
	}
	bsint := make(map[int]int)
	for _, x := range bslice {
		bsint[x]++
	}
	for idx, x := range bsint {
		if x == 1 {
			fmt.Printf(" число %d попалось 1 раз \n", idx)
		}
	}

}
