package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return r.height*2 + r.width*2
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	rect := rectangle{height: 5, width: 10}
	fmt.Println(rect.area())

	geo := geometry(rect)
	fmt.Println(geo.area())

	circ := circle{radius: 10}
	geo = geometry(circ)

}
