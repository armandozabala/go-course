package main

import "fmt"

// variables globales fuera de las funciones
var mensaje string

func funcion1() {
	mensaje = "Modificada"
	fmt.Println(mensaje)
}

func main() {
	mensaje = "Go run"
	fmt.Println(mensaje)

	//ejecuta cuando termina de ejecutar main con el uso del defer
	defer funcion1()

	fmt.Println(mensaje)
}
