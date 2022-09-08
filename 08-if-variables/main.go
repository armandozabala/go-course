package main

import "fmt"

func main() {

	//declarar el nombre en el if
	if nombre, edad := "Alex", 26; nombre == "Alex" {
		fmt.Println("Hola ", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d \n", nombre, edad)
	}

	//Crear mapa
	users := make(map[string]string)

	users["alex"] = "alex@gmail.com"
	users["roel"] = "roel@gmail.com"

	//vacio por que no existe
	fmt.Println(users["juan"])

	//correo, ok := users["alex"]

	if correo, ok := users["alex"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible")
	}

}
