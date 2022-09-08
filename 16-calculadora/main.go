package main

import (
	"fmt"
	"strconv"
	"strings"
)

//manejo de errores

func sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	var result int

	for _, valor := range valores {
		//convertir string a entero
		num, error := strconv.Atoi(valor)
		if error != nil {
			fmt.Println(error)
			fmt.Println("Tiene un error numero enteros")
			fmt.Println("Solo debe usar el operador mas")
		} else {
			result += num
		}

	}
	return result
}

func main() {

	var expresion string
	var result int

	fmt.Print("=>")
	fmt.Scanln(&expresion)

	result = sumar(expresion)

	fmt.Printf("=> %d \n", result)
}
