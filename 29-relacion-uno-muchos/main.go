package main

import "fmt"

// Relacion de uno a muchos
type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {

	//Relacion de uno a muchos
	v1 := Video{
		titulo: "video GO 1",
	}
	v2 := Video{
		titulo: "Video GO 2",
	}

	curso := Curso{
		titulo: "Curso de Go",
		videos: []Video{
			v1,
			v2,
		},
	}

	v1.curso = curso
	v2.curso = curso

	fmt.Println(v1.curso.titulo)

	//For para correr los videos
	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}

}
