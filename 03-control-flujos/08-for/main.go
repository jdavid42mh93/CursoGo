package main

import "fmt"

func main() {
	//bucle infinito
	/*for {
		fmt.Println("Hola Mundo")
	}*/

	//bucle tipo while
	numeros := 12455
	c := 0
	for numeros > 0 {
		numeros /= 10
		c++
	}
	fmt.Println("Cantidad de digitos es", c)

	//bucle for

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("Hola Mundo", i)
		}
	}
}
