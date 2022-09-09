package main

import "fmt"

// Interfaces
type Animal interface {
	mover() string
}

type Perro struct{}
type Pez struct{}
type Pajaro struct{}

// Interface maneja los metodos identicos de diferentes estructuras
func (*Perro) mover() string {
	return "Soy Perro y camino"
}

func (*Pajaro) mover() string {
	return "Soy Pajaro y vuelo"
}

func (*Pez) mover() string {
	return "Soy Pez y nado"
}

func moverAnimal(animal Animal) {
	fmt.Println(animal.mover())
}

func main() {

	perro := Perro{}
	moverAnimal(&perro)

	pez := Pez{}
	moverAnimal(&pez)

	pajaro := Pajaro{}
	moverAnimal(&pajaro)

}
