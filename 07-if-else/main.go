package main

import "fmt"

func main() {

	var consumo, descuento float64
	var datoDescuento string

	//Entrada
	fmt.Print("Ingrese el consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		//Descuento de 10%
		if consumo <= 100 {
			datoDescuento = "10%"
			descuento = consumo * 0.10
		} else if consumo > 100 && consumo <= 200 {
			datoDescuento = "20%"
			descuento = consumo * 0.20
		} else if consumo > 200 {
			datoDescuento = "30%"
			descuento = consumo * 0.30
		}

	} else {
		fmt.Println("Error al ingresar el consumo")
	}

	//Operaciones
	montoDescuento := consumo - descuento
	igv := montoDescuento * 0.19
	total := montoDescuento + igv

	fmt.Println("-------FACTURA-------")
	fmt.Println("Descuento que aplica (", datoDescuento, ")", descuento)
	fmt.Println("Consumo: ", consumo)
	fmt.Println("Monto de descuento: ", montoDescuento)
	fmt.Println("Igv: ", igv)
	fmt.Println("Total a Pagar: ", total)

}
