package main

import "fmt"

func main() {
	var consumo, descuento float64
	var datosDescuento string

	fmt.Print("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		if consumo <= 100 {
			datosDescuento = "10%"
			descuento = consumo * 0.10
		} else if consumo > 100 && consumo <= 200 {
			datosDescuento = "20%"
			descuento = consumo * 0.20
		} else if consumo > 200 {
			datosDescuento = "30%"
			descuento = consumo * 0.30
		}
	} else {
		fmt.Println("Error al ingresar el consumo")
	}

	montoDescuento := consumo - descuento
	igv := montoDescuento * 0.12
	totalPagar := montoDescuento + igv

	fmt.Println("--------- FACTURA DE CONSUMO ---------")
	fmt.Println("Descuento que aplica", datosDescuento)
	fmt.Println("Consumo", consumo)
	fmt.Println("Descuento", descuento)
	fmt.Println("Monto con descuento", montoDescuento)
	fmt.Println("IGV", igv)
	fmt.Println("Total a pagar", totalPagar)
}
