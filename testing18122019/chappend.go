package main

import "fmt"

func main() {
	var runes []rune
	for _, r := range "Hello my дорогой Всилий" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

}
