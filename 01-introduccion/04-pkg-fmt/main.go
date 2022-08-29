package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo!!"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Juan"
	edad := 28

	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)

	fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)

	message := fmt.Sprintf("Hola, %v y su edad es %v", nombre, edad)

	fmt.Println(message)

	fmt.Printf("nombre: %T \n", edad)

	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("otro nombre: ", nombre)
}
