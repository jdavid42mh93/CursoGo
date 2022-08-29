package main

import "fmt"

func main() {
	dias := make(map[int]string)

	fmt.Println(dias)

	//agregar datos
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"

	fmt.Println(dias)

	//modificar datos
	dias[7] = "SABADO"
	fmt.Println(dias)

	//eliminar datos
	delete(dias, 1)
	fmt.Println(dias, len(dias))

	//nuevo map
	estudiantes := make(map[string][]int)

	estudiantes["Juan"] = []int{13, 15, 16}
	estudiantes["David"] = []int{18, 17, 15}

	fmt.Println(estudiantes)

	fmt.Println(estudiantes["Juan"][2])
}
