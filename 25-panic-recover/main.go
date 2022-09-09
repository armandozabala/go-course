package main

import (
	"fmt"
	"os"
)

func main() {

	//Funcion
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Ups al parecer no finalizo de forma correcta")
		}
	}()

	// leer archivo
	if file, error := os.Open("hoa.txt"); error != nil {
		//para enviar una excepcion y detener el programa
		panic("No fue posible")
		//funcion recove controla los panic

	} else {
		//defer se va ejecutar al final de la ejecucion, por que no podemos poner close al final
		defer func() {
			fmt.Println("Cerrando")
			file.Close()
		}()
		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)

		contenidoArchivo := string(contenido[:long])

		fmt.Println(contenidoArchivo)
	}

}
