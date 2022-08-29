package main

import "fmt"

func main() {
	var numeros [5]int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	fmt.Println(numeros)

	nombres := [3]string{"Juan", "David", "Gabriela"}
	fmt.Println(nombres)

	colores := [...]string{
		"Rojo",
		"Verde",
		"Negro",
		"Azul",
	}

	fmt.Println(colores, len(colores))

	//indices definidos
	monedas := [...]string{
		0: "Dolar",
		2: "Euro",
		3: "Soles",
		5: "Pesos",
	}

	fmt.Println(monedas, len(monedas))

	//subarreglo
	submoneda := monedas[0:3]
	fmt.Println(submoneda)
}
