package main

import (
	"fmt"
	m "shop/shopcompetition"
)

func main() {
	yList := m.AccountList{}
	nameList := []string{"Vasily", "Petr", "Oleg"}
	for _, name := range nameList {
		yList.Register(name)
	}

	fmt.Println(yList)

	for _, name1 := range nameList {
		yList.AddBalance(name1, 100)
	}

	fmt.Println(yList)
}
