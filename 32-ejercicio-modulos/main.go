package main

import "paquetes/figuras"

func main() {
	cua1 := figuras.Cuadrado{Lado: 8}
	figuras.Medidas(&cua1)
}
