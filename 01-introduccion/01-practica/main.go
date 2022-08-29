package main

import "fmt"

// Practica 01: Suma de dos números enteros
func sumar(a, b int) int {
	return a + b
}

// Practica 02: Calcular cociente y el Residuo de dos números enteros
func cociente(a, b int) int {
	return a / b
}

func residuo(a, b int) int {
	return a % b
}

// *****************************************************************

// Practica 03: Calcular el Precio de Venta
func precioVenta(vv, igv float64) float64 {
	return igv + vv
}

func calcularIGV(vv, igv float64) float64 {
	return igv * vv
}

func main() {
	var resultado, a, b int
	var vv, igv, pv float64

	fmt.Println("Ingrese un numero A: ")
	fmt.Scanln(&a)
	fmt.Println("Ingrese un numero B: ")
	fmt.Scanln(&b)

	resultado = sumar(a, b)
	fmt.Printf("El resultao de la suma es %d \n", resultado)
	resultado = cociente(a, b)
	fmt.Printf("El cociente es %d \n", resultado)
	resultado = residuo(a, b)
	fmt.Printf("El residuo es %d \n", resultado)

	fmt.Println("Ingrese el valor de venta: ")
	fmt.Scanln(&vv)

	//calcular igv
	igv = calcularIGV(vv, 0.18)

	//calcular precio de venta
	pv = precioVenta(vv, igv)

	//Salida de datos
	fmt.Println("==========================")
	fmt.Println("Valor de Venta:", vv)
	fmt.Println("IGV:", igv)
	fmt.Println("Precio de Venta:", pv)
	fmt.Println("==========================")
}
