package main

import "fmt"

func saludar(nombre string){
	//fmt.Println("Saludos desde la funcion");
	fmt.Println("Hola",nombre);
}

func sumar(a, b int) int {
   return a + b;
}


func main(){
	saludar("Armando");
	res:=sumar(10,20);
	fmt.Println("La Suma es: ",res);
}