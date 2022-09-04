package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {
	/* relacion uno a uno
	juan := User{
		nombre: "Juan",
		email:  "jdavid42mh93@gmail.com",
		activo: true,
	}

	davidStudent := Student{
		user:   juan,
		codigo: "jdavid",
	}
	fmt.Println(juan)
	fmt.Println(davidStudent.user.nombre)*/

	/* relacion de uno a muchos */
	video1 := Video{titulo: "01-introudccion"}
	video2 := Video{titulo: "02-instalacion"}

	curso := Curso{
		titulo: "Curso de Go",
		videos: []Video{video1, video2},
	}

	video1.curso = curso
	video2.curso = curso
	fmt.Println(video1.curso.titulo)

	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}

}
