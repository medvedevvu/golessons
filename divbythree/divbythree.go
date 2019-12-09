package main

import "fmt"

func d3() {
	for i := 1; i <= 100; i += 1 {
		if i%3 == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}

func FizzBuzz() {
	for i := 1; i <= 100; i += 1 {
		switch {
		case i%3 == 0 && !(i%5 == 0):
			{
				fmt.Printf("%s\n", "Fizz")
			}
		case i%5 == 0 && !(i%3 == 0):
			{
				fmt.Printf("%s\n", "Buzz")
			}
		case (i%5 == 0) && (i%3 == 0):
			{
				fmt.Printf("%s\n", "FizzBuzz")
			}
		default:
			fmt.Printf("%d\n", i)
		}
	}
}
func main() {
	d3()
	FizzBuzz()
}
