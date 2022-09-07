package main

import "fmt"

func main(){
	a:=20;
	b:=10;

	result := a + b;

	fmt.Println("Suma: ", result)

	result = a - b;

	fmt.Println("Resta: ",result);

	result = a * b;

	fmt.Println("Mult: ",result);

	result = a / b;

	fmt.Println("Division: ", result);

	var divi float64= 3.0 / 2.0;

	fmt.Println("Division: ", divi);

	result = 3 % 2;

	fmt.Println("Modulo: ",result);
}