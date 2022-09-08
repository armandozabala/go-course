package main

import "fmt"

//Recursividad

func factorial(n int) int {

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {

	//factorial
	fmt.Println(factorial(3))
}
