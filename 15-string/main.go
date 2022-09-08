package main

import (
	"fmt"
	"strings"
)

func esPalindromo(palabra string) {
	fmt.Println(palabra)
	//palabra to Lower
	palabra = strings.ToLower(palabra)

	fmt.Println(palabra)
}

func main() {

	esPalindromo("Oso")

}
