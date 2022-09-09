package main

import "fmt"

// Struc
type Persona struct {
	nombre string
	edad   int
}

// Definir metodos, se envia la referencia por eso se usa *
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s \n Edad: %d \n", p.nombre, p.edad)
}

// Metodo para modificar el nombre del atributo del objeto
func (p *Persona) editNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

// Herencia de la otra estrcutura de Persona
type Empleado struct {
	Persona
	sueldo float64
}

// Crear objeto de la estructura
func main() {

	persona1 := Persona{"Armando", 36}

	persona1.imprimir()
	//fmt.Println(persona1)

	// se puede modificar valor del objeto
	//persona1.nombre = "Jose"
	//persona1.edad = 40
	persona1.editNombre("Juancho")
	persona1.imprimir()
	//fmt.Println(persona1)

	p2 := Persona{
		nombre: "Pedro",
		edad:   25,
	}
	p2.imprimir()

	p2.editNombre("Perucho")
	p2.imprimir()

	//fmt.Println(p2)

	// nnuevo objeto con la herencia
	em1 := Empleado{
		sueldo: 30.99,
	}
	em1.nombre = "Antonio"
	em1.edad = 30
	em1.imprimir()

	fmt.Println(em1)

}
