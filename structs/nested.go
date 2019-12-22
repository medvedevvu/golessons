package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
}

type Whell struct {
	Circle
	Spokes int
}

func main() {
	var w Whell
	w.X = 10
	w.Y = 100
	w.Radius = 5
	w.Spokes = 20

	fmt.Printf("%#v\n", w)

}
