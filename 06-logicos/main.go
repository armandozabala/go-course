package main

import "fmt"

func main() {

	//Operadores Logico

	//Not
	fmt.Println(!true)
	//And
	fmt.Println(true && true)
	//Or
	fmt.Println(false || true)

	b := 2

	r := b == 2 && !(4 > b)

	fmt.Println("R false", r)

	res := b == 2 && (4 > b)

	fmt.Println("R true", res)
}
