package main

import "fmt"

func main() {
	nombres := [...]string{"Enero", "Febrero", "Marzo"}

	//con el indice
	for indice, elementos := range nombres {
		fmt.Println(indice, elementos)
	}

	//sin los indices se coloca _
	for _, elementos := range nombres {
		fmt.Println(elementos)
	}

}
