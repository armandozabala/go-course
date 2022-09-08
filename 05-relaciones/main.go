package main

import "fmt"

func main(){

	a := 3
	b := 4

	var r bool
	
				//igualdad
				r = a == b
				fmt.Printf("%d es igual que %d? %t \n",a,b,r)

				//distinto
				r = a != b
				fmt.Printf("%d es es distinto que %d? %t \n",a,b,r)


				//mayor
				r = a > b
				fmt.Printf("%d es mayor que %d? %t \n",a,b,r)

				//menor
				r = a < b
				fmt.Printf("%d es menor que %d? %t \n",a,b,r)
				
				//mayor o igual que
				r = a >= b
				fmt.Printf("%d es mayor o igual que %d? %t \n",a,b,r)

				//menor o igual que
				r = a <= b
				fmt.Printf("%d es menor o igual que %d? %t \n",a,b,r)


}