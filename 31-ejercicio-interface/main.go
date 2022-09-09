package main

import (
	"fmt"
	"math"
)

type Geometrica interface {
	area() float32
	perimetro() float32
}

type Cuadrado struct {
	lado float32
}
type Circulo struct {
	radio float32
}

func (c *Cuadrado) area() float32 {
	return c.lado * c.lado
}

func (c *Cuadrado) perimetro() float32 {
	return c.lado * 4
}

func (c *Circulo) area() float32 {
	return math.Pi * (c.radio * c.radio)
}

func medidas(g Geometrica) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {

	cua1 := Cuadrado{lado: 10}

	medidas(&cua1)
}
