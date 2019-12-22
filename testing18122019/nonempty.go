package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func remove(slice []int, i int) []int {
	fmt.Println(slice)
	copy(slice[i:], slice[i+1:])
	fmt.Println(slice)
	return slice[:len(slice)-1]
}

func main() {
	tmp := []string{"ass", "ddd", "", "x"}
	fmt.Println(nonempty(tmp))
	var smap2 map[string]string
	smap := map[string]map[int]string{
		"xxx": map[int]string{1: "s", 2: "d"}}
	fmt.Println(smap)
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
	fmt.Println(smap2 == nil)
}
