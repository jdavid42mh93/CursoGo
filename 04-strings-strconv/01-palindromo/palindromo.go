package main

import (
	"fmt"
	"strings"
)

func reverse(cadena string) string {
	arrayCadena := strings.Split(cadena, "")
	arrayInvertida := make([]string, 0)

	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[i])
	}
	return strings.Join(arrayInvertida, "")
}

func esPalindromo(palabra string) bool {
	palabra = strings.ToLower(palabra)
	//palabra = strings.Replace(palabra, "z", "s", 2)
	palabra = strings.ReplaceAll(palabra, " ", "")
	palabraInvertida := reverse(palabra)
	return palabra == palabraInvertida
}
func main() {
	esPalindromo("Juan")
	if esPalindromo("Luz azul") {
		fmt.Println("es palindromo")
	} else {
		fmt.Println("no es palindromo")
	}
}
