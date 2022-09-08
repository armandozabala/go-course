package main

import "fmt"

func main() {

	//funcion anonimas
	/*
		func() {
			fmt.Println("Hola desde la funcion anonimas")
		}()*/

	myfunc := func(nombre string) string {
		return fmt.Sprintf("Hola, %s desde funcion con variable", nombre)
	}

	fmt.Println(myfunc("Armando"))
}
