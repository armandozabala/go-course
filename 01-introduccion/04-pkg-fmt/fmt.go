package main

import "fmt"

func main(){

	hola := "Hola";
	mundo := "Mundo";

	fmt.Println(hola);
	fmt.Println(mundo);

	nombre := "Armando";
	edad := 36;

	fmt.Printf("* Hola, %s y su edad es %d \n",nombre, edad);

	fmt.Printf("* Hola, %v y su edad es %v \n",nombre, edad);

	mensaje := fmt.Sprintf("Hola, %v y su edad es %v",nombre, edad);

	fmt.Println(mensaje);

	fmt.Printf("nombre %T \n", nombre);
	fmt.Printf("edad %T \n", edad);


	fmt.Print("Ingrese su nombre: ");
	fmt.Scanln(&nombre);

	fmt.Println("Otro nombre es: ",nombre);
}