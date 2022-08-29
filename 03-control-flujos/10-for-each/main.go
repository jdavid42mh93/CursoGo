package main

import "fmt"

func main() {
	nombres := [...]string{"Juan", "David", "Evelyn", "Andrea"}

	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])
	}

	for indice, elemento := range nombres {
		fmt.Println(indice, elemento)
	}

	for _, elemento := range nombres {
		fmt.Println(elemento)
	}

	for indice, _ := range nombres {
		fmt.Println(indice)
	}
}
