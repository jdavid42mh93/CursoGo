package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hola ", nombre)
}

func sumar(a, b int) int {
	return a + b
}
func main() {
	var nombre string
	var resultado, a, b int
	fmt.Println("Ingrese un nombre")
	fmt.Scanln(&nombre)
	saludar(nombre)

	fmt.Println("Ingrese un numero A: ")
	fmt.Scanln(&a)
	fmt.Println("Ingrese un numero B: ")
	fmt.Scanln(&b)

	resultado = sumar(a, b)
	fmt.Printf("El resultao de la suma es %d \n", resultado)

}
