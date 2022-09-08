package main

import "fmt"

// funcion varidica, recibe N numeros
// (string, int) asi se coloca para multiple retorno
func sumar(nombre string, numeros ...int) (string, int) {

	mensaje := fmt.Sprintf("La suma de %s es: ", nombre)
	fmt.Println(numeros)
	var total int
	for _, num := range numeros {
		total += num
	}

	//multiple valores para el retorno
	return mensaje, total
}

func main() {

	mensaje, result := sumar("Armando", 4, 4, 5, 6, 20, 10)

	fmt.Print(mensaje)
	fmt.Println(result)
}
