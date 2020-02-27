package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	x,y float64
}

func (v Vertex) abs() float64{
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func (v *Vertex) Scale(n float64){
	v.x *= n
	v.y *= n
}

func main() {
	v := Vertex{3,4}
	fmt.Println(v.abs())

	v.Scale(10)
	fmt.Println(v.abs())
}
