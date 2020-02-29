package main

import "fmt"

type Point struct {
	x,y float64
}

func main() {
	var i interface{}
	describe(i)

	i = 25
	describe(i)

	i = Point{2.5, 6.4}
	describe(i)

}

func describe( v interface{}){
	fmt.Printf("%v  %T\n", v , v)
}