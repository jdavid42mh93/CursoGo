package main

import "fmt"

func contadorPN() {
	var num, positivo, negativo int
	fmt.Print("Ingrese un valor: ")
	fmt.Scan(&num)
	for num != 0 {
		if num > 0 {
			positivo += 1
		} else {
			negativo += 1
		}
		fmt.Print("Ingrese un valor: ")
		fmt.Scan(&num)
	}
	fmt.Println("# Positivos", positivo)
	fmt.Println("# Negativos", negativo)
	fmt.Println("Fin del programa")
}

func contadorDigitos(arreglo [10]int) {
	var digito1, digito2, digito3 int
	for _, elemento := range arreglo {
		if elemento < 10 {
			digito1++
		} else {
			if elemento < 100 {
				digito2++
			} else {
				digito3++
			}
		}
	}
	fmt.Println(digito1, digito2, digito3)
}

func main() {
	//contadorPN()
	vector := [10]int{100, 30, 1, 4, 6, 2, 43, 101, 300, 3}
	contadorDigitos(vector)
}
