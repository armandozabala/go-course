package main

import (
	"fmt"
	"os"
)

func main() {
	// leer archivo
	// creamos variables dentro del if
	if file, error := os.Open("hola.txt"); error != nil {
		fmt.Println("No fue posible")
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
