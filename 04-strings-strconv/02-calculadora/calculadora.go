package main

import (
	"fmt"
	"strconv"
	"strings"
)

func suma(expresion string) int {
	valores := strings.Split(expresion, "+")
	var result int
	for _, valor := range valores {
		num, error := strconv.Atoi(valor)
		if error != nil {
			fmt.Println(error)
			fmt.Println("Error: Tiene que ingresar un numero entero")
			fmt.Println("o Solo debes realizar con operador +!!")
		} else {
			result += num
		}
		fmt.Println(num, error)
	}
	return result
}

func main() {
	var expresion string
	var result int

	fmt.Print("=>")
	fmt.Scanln(&expresion)
	result = suma(expresion)

	fmt.Println(result)
}
