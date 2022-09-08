package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {

		if i == 5 {
			fmt.Println("Salto en ")
			//salta a la siguiente iteraccion
			continue
		}
		if i == 8 {
			fmt.Println("Vamos a romper")
			break
		}
		fmt.Println(i)
	}

}
