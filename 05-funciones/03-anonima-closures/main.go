package main

import (
	"fmt"
	"strings"
)

// closures : funcion que retorna una funcion
func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {
	// func anonima
	/* func() {
		fmt.Println("Hola desde la funcion anonima")
	}()

	myfunc := func(nombre string) string {
		return fmt.Sprintf("Hola, %s desde la funcion anonima", nombre)
	}
	fmt.Println(myfunc("Juan"))*/
	repeat3 := repeat(4)

	fmt.Println(repeat3("Juan David "))
}
