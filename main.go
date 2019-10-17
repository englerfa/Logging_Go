package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func sayHello(s string) {
	fmt.Println("hello " + s)
}

func main() {
	//v := Vertex{3, 4}
	//fmt.Print("Abs() is ")
	//fmt.Println(v.Abs())

	sayHello("world")
}
