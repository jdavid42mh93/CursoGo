package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			fmt.Println("saltar la iteracion")
			continue
		}

		if i == 8 {
			fmt.Println("romper el bucle")
			break
		}
		fmt.Println(i)
	}
}
