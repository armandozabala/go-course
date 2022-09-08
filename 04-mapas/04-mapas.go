package main

import "fmt"

func main(){ 

	//Mapas
	//como los objetos
	
	dias := make(map[int]string)

	fmt.Println(dias);

	//Agregar datos
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	
	fmt.Println(dias)

	dias[4]="Sabado"
	//listar
	fmt.Println(dias)

	//Eliminar
	delete(dias, 1);

	//Mostrar Logintu
	fmt.Println(dias, len(dias))


}