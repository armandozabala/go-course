package main

import "fmt"

func main(){
	
	var nombre1 string;
	nombre1="Armando";
	var nombre2 string = "Jose";
	edad:=36;
	promedio:= 25.5;


	fmt.Println(nombre1, nombre2, edad, " Promedio ", promedio);

	var a int // valor 0 por defecto
	var b float64 // valor 0
	var c string // vacio
	var d bool // valor por defecto false

	fmt.Println(a,b,c,d);


	const pi = 3.141592;

	fmt.Println(pi);


}