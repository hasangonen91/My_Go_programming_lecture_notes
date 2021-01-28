package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	//
	//birden fazla ozellikte olabilir
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

// bu shape area s.area aslında area imzali her seyi cagir demek
func school(s shape) {
	fmt.Println("Seklin alani : ", s.area())
}

/*
func Demo1() {
	fmt.Println("yukseklik girin: ")
	var yukseklik int
	fmt.Scanln(&yukseklik)
	fmt.Println("genislik girin")
	var genislik int
	fmt.Scanln(&genislik)
	r := rectangle{width: float64(genislik), height: float64(yukseklik)}
	school(r)
}*/
func Demo1() {
	r := rectangle{width: 10, height: 6}
	school(r)
	c := circle{radius: 6.86}
	school(c)
}
