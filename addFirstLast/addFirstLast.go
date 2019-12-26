// задания 2 , 3
package main

import "fmt"

/* 2,3 */
func main() {
	myslice := ([]int{7, 7, 7, 7, 7, 7, 7, 7})[:]
	tmp := make([]int, len(myslice)+1)
	tmp[0] = 5
	copy(tmp[1:], myslice) // в переди
	fmt.Println(tmp)
	fmt.Println(append(myslice, 5)) // позади
}
