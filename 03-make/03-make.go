package main

import "fmt"

func main() {
	//Make
	numeros := make([]int, 3, 3)

	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300

	fmt.Println(numeros, len(numeros), cap(numeros))

}