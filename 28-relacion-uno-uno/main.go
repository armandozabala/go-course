package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

// se crea una variable y se relacion de uno a uno
type Student struct {
	user   User
	codigo string
}

func main() {

	alex := User{
		nombre: "Alex",
		email:  "a@gmail.com",
		activo: true,
	}

	roel := User{
		nombre: "Roel",
		email:  "r@gmail.com",
		activo: true,
	}

	alexStudent := Student{
		user:   alex,
		codigo: "3123",
	}

	fmt.Println(roel)
	fmt.Println(alexStudent)
}
