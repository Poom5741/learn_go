package main

import "fmt"

func main() {
	var a interface{}

	a = 10
	fmt.Printf("%T %v\n", a, a)

	a = "ten"
	fmt.Printf("%T %v\n", a, a)

	a = true
	fmt.Printf("%T %v\n", a, a)

	a = func() string { return "hello" }
	fmt.Printf("%T %v", a, a)
}

type volumer interface {
	Volume() float64
}

type cube struct {
	edge float64
} // edge x edge x edge

type triangularPrism struct {
	base     float64
	attitude float64
	height   float64
} // 0.5 x base x attitude x height

func VolumeOf(v volumer) float64 {
	v = triangularPrism.base * triangularPrism.attitude * triangularPrism.height
	return v
}
