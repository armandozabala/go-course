package main

import "fmt"

func main() {
	//Bucle For
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("Hola ", i)
		}
	}

	//Infinito
	/*for {
		fmt.Println("Hola Mundo")
	}*/

	//Bucle tipo While
	/*numero := 12334
	c := 0
	for numero > 0 {
		numero /= 10
		c++
	}*/

	//fmt.Println("Cantidad de digito ", c)
}
