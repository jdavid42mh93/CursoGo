package main

import "fmt"

func notas() {
	var nota int
	conta1 := 0
	conta2 := 0
	x := 1
	for x <= 10 {
		fmt.Print("Ingrese la nota del alumno: ")
		fmt.Scan(&nota)
		if nota > 7 {
			conta1 += 1
		} else {
			conta2 += 1
		}
		x += 1
	}
	fmt.Println("Cantidad de alumnos con notas mayores o iguales a 7:", conta1)
	fmt.Println("Cantidad de alumnos con notas menores a 7:", conta2)
}

func alturas() {
	var npersonas int32
	var altura, suma, promedio float32
	fmt.Print("Ingrese el numero de personas: ")
	fmt.Scan(&npersonas)
	suma = 0
	i := 0
	for i < int(npersonas) {
		fmt.Print("Ingrese la altura de la persona: ")
		fmt.Scan(&altura)
		suma = suma + altura
		i += 1
	}
	promedio = suma / float32(npersonas)
	fmt.Println("Altura promedio es:", promedio)
}
func empleados() {
	var nempleados int32
	var sueldo float32
	fmt.Print("ingrese el numero de empleados: ")
	fmt.Scan(&nempleados)
	conta1 := 0
	conta2 := 0
	suma := 0
	i := 0
	for i < int(nempleados) {
		fmt.Print("ingrese el sueldo del empleado: ")
		fmt.Scan(&sueldo)
		if sueldo < 100 || sueldo > 500 {
			fmt.Println("Sueldo no se encuentra dentro del rango")
		} else {
			if sueldo >= 100 && sueldo <= 300 {
				conta1 += 1
			} else if sueldo > 300 {
				conta2 += 1
			}
			suma = suma + int(sueldo)
		}
		i += 1
	}
	fmt.Println("Sueldo al personal:", suma)
	fmt.Println("Cantidad de persona con sueldo entre 100 y 300: ", conta1)
	fmt.Println("Cantidad de persona con sueldo mayor a 300: ", conta2)
}

func serie() {
	i := 0
	n := 11
	for i < 25 {
		n += 11
		fmt.Print(n, "-")
		i += 1
	}
	fmt.Print("\n")
}

func multiplo8() {
	n := 8
	for n < 500 {
		fmt.Print(n, "-")
		n += 8
	}
	fmt.Print("\n")
}
func lista1() int32 {
	var n, suma int32
	i := 0
	suma = 0
	for i < 15 {
		fmt.Print("Ingrese un valor: ")
		fmt.Scan(&n)
		i += 1
		suma = suma + n
	}
	fmt.Println(suma)
	return suma
}

func lista2() int32 {
	var n, suma int32
	i := 0
	suma = 0
	for i < 15 {
		fmt.Print("Ingrese un valor: ")
		fmt.Scan(&n)
		i += 1
		suma = suma + n
	}
	fmt.Println(suma)
	return suma
}

func main() {
	//notas()
	//alturas()
	//empleados()
	//serie()
	//multiplo8()
	var suma1, suma2 int32
	fmt.Println("Datos de la Lista 1")
	suma1 = lista1()
	fmt.Println("Datos de la Lista 2")
	suma2 = lista2()
	if suma1 == suma2 {
		fmt.Println("Lista iguales: ", suma1, suma2)
	} else {
		if suma1 > suma2 {
			fmt.Println("Lista 1 mayor: ", suma1, suma2)
		} else {
			fmt.Println("Lista 2 mayor: ", suma2, suma1)
		}
	}
}
