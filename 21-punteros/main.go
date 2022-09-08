package main

import "fmt"

// punteros
func modificarValor(cadena *string) {
	//ya se recibe la referencia
	fmt.Printf("%p \n", cadena)
	*cadena = "Hola desde la funcion"
}

func main() {

	cadena := "Hola Mundo de Go"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la funcion: ", cadena)

	modificarValor(&cadena) // se pasa la referencia

	fmt.Println("Despues de la funcion: ", cadena)
}
