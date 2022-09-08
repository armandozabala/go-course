package main

import "fmt"

func main(){
	//Slicen
	numeros :=  []int{1,2,3}

	fmt.Println(numeros, len(numeros));

	//Agregar datos
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)

	fmt.Println(numeros, len(numeros))

	//Sub slicen
	subSlicen := numeros[:2]
	//referencia 
	numeros[0] = 100

	fmt.Println(subSlicen)
	fmt.Println(numeros)

	//Caracteristica de Slicen

	meses :=  []string{"Enero","Febrero","Marzo"}

	fmt.Printf("Longitud: %v, Cap: %v, %p \n",len(meses), cap(meses), meses)

	meses = append(meses, "Abril");

	fmt.Printf("Longitud: %v, Cap: %v, %p \n",len(meses), cap(meses), meses)

}