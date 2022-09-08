package main

import (
	"fmt"
	"strings"
)

// Closures
// repeat recibe un numero entero
// retorna otra funcion que rcibe un string y retorna un string
func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {
	// funcion closures
	// tiene la referencia de la funcion padre
	repeat3 := repeat(3)
	//recuerda el valor que se envio inicialmente
	fmt.Println(repeat3("Armando"))

	repeat5 := repeat(5)

	fmt.Println(repeat5("Zabala"))

}
