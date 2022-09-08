package main

import (
	"fmt"
	"strings"
)

func reverse(palabra string) string {
	arrayCadena := strings.Split(palabra, "")
	arrayInvertida := make([]string, 0)

	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[i])
	}

	//convierte en string
	return strings.Join(arrayInvertida, "")
}

func esPalindromo(palabra string) bool {
	//opalabra original
	fmt.Println(palabra)
	//palabra to Lower
	palabra = strings.ToLower(palabra)
	palabra = strings.Replace(palabra, "z", "s", 2)
	palabra = strings.ReplaceAll(palabra, " ", "")

	palabraInvertida := reverse(palabra)

	return palabra == palabraInvertida
}

func main() {

	if esPalindromo("Luz") {
		fmt.Println("Es Palindrome")
	} else {
		fmt.Println("No Es Palindrome")
	}

}
