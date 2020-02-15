package main

import (
	"fmt"
	sc "shop_competition"
)

func main() {
	acc := sc.NewAccountsList()
	err := acc.Register("Lova")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(acc)
}
