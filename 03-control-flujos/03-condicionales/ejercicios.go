package main

import "fmt"

func operaciones() {
	var numero1, numero2 int32
	fmt.Println("Ingrese el primer numero:")
	fmt.Scan(&numero1)

	fmt.Println("Ingrese el segundo numero:")
	fmt.Scan(&numero2)

	if numero1 > numero2 {
		fmt.Println("La suma de los numeros es: ", numero1+numero2)
		fmt.Println("La resta de los numeros es: ", numero1-numero2)
	} else {
		fmt.Println("La multiplicacion de los numeros es: ", numero1*numero2)
		if numero2 == 0 {
			fmt.Println("No es posible la division por 0")
		} else {
			fmt.Println("La division de los numeros es: ", numero1/numero2)
		}
	}
}

func promedio() {
	var nota1, nota2, nota3, promedio float32
	fmt.Println("Ingrese la primera nota:")
	fmt.Scan(&nota1)

	fmt.Println("Ingrese la segunda nota:")
	fmt.Scan(&nota2)

	fmt.Println("Ingrese la tercera nota:")
	fmt.Scan(&nota3)

	promedio = (nota1 + nota2 + nota3) / 3
	if promedio >= 7.0 {
		fmt.Println("Promocionado")
	} else {
		fmt.Println("Reprobado")
	}
}

func positivo() {
	var numero1 int32
	fmt.Println("Ingrese el primer numero:")
	fmt.Scan(&numero1)

	if numero1 >= 100 {
		fmt.Println("El numero ingreado tiene tres digitos")
	} else if numero1 < 10 {
		fmt.Println("el numero ingresado tiene un digito")
	} else {
		fmt.Println("el numero ingresado tiene dos digitos")
	}
}

func promedio2() {
	var nota1, nota2, nota3, promedio float32
	fmt.Println("Ingrese la primera nota:")
	fmt.Scan(&nota1)

	fmt.Println("Ingrese la segunda nota:")
	fmt.Scan(&nota2)

	fmt.Println("Ingrese la tercera nota:")
	fmt.Scan(&nota3)

	promedio = (nota1 + nota2 + nota3) / 3
	if promedio >= 7.0 {
		fmt.Println("Promocionado")
	} else if promedio >= 4 && promedio < 7 {
		fmt.Println("Regular")
	} else if promedio < 4 {
		fmt.Println("Reprobado")
	}
}

func mayor() {
	var nota1, nota2, nota3, mayor int32
	fmt.Println("Ingrese la primera nota:")
	fmt.Scan(&nota1)

	fmt.Println("Ingrese la segunda nota:")
	fmt.Scan(&nota2)

	fmt.Println("Ingrese la tercera nota:")
	fmt.Scan(&nota3)

	if nota1 > nota2 {
		mayor = nota1
	} else {
		mayor = nota2
	}
	if mayor < nota3 {
		mayor = nota3
	}
	fmt.Println("el mayor es:", mayor)
}

func trescifras() {
	var numero int32
	fmt.Println("Ingrese el primer numero:")
	fmt.Scan(&numero)

	if numero < 0 {
		fmt.Println("Debe ser un numero positivo")
	}
	if numero >= 0 && numero < 10 {
		fmt.Println("el numero tiene un digito")
	} else if numero >= 10 && numero < 100 {
		fmt.Println("el numero tiene dos digitos")
	} else if numero >= 100 && numero < 1000 {
		fmt.Println("el numero tiene tres cifras")
	} else if numero >= 1000 {
		fmt.Println("Debe ser un numero valido")
	}
}

func capacitacion() {
	var totalPreguntas, preguntasContestadas int32
	var porcentaje float32
	fmt.Println("Ingrese el total de preguntas:")
	fmt.Scan(&totalPreguntas)

	fmt.Println("Ingrese el numero de preguntas contestadas:")
	fmt.Scan(&preguntasContestadas)

	porcentaje = (float32(preguntasContestadas) * 100) / float32(totalPreguntas)

	if porcentaje >= 90 {
		fmt.Println("Nivel maximo")
	} else if porcentaje >= 75 && porcentaje < 90 {
		fmt.Println("Nivel medio")
	} else if porcentaje >= 50 && porcentaje < 75 {
		fmt.Println("Nivel regular")
	} else if porcentaje < 50 {
		fmt.Println("Fuera de nivel")
	}
}

func main() {
	//operaciones()
	//promedio()
	//positivo()
	//promedio2()
	//mayor()
	//trescifras()
	capacitacion()
}
