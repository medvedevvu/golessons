// задание 1
package main

import (
	"fmt"
	"strings"
)

func main() {
	var lText string = "спички топор молоток к веревка спички удочка молоток "
	lTextarray := strings.Split(lText, " ")
	lTextmap := make(map[string]int)
	for _, elem := range lTextarray {
		if len(strings.Trim(elem, " ")) > 0 {
			lTextmap[elem]++
		}
	}
	for idx, val := range lTextmap {
		fmt.Printf(" слово \"%s\" попалось %d раз \n", idx, val)
	}

}
