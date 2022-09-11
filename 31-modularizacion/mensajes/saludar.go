package mensajes

import "fmt"

func Hola() {
	fmt.Println("Saludos desde mensajes")
}

const mensaje = "Hola desde constante"

// funcion publica inicia con Mayuscula
func Imprimir() {
	fmt.Println(mensaje)
	funcionPrivada()
}

// funcion privada se usa solo
func funcionPrivada() {
	fmt.Println("Hola desde la funcion privada")
}
