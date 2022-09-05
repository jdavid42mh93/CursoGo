package main

import (
	"fmt"
	"paquetes/models"

	"github.com/donvito/hellomod"
)

func main() {
	/*
		mensajes.Hola()
		mensajes.Impirmir()
	*/

	/*
		cuadrado := figuras.Cuadrado{Lado: 5}
		figuras.Medidas(&cuadrado)

		circulo := figuras.Circulo{Radio: 3}
		figuras.Medidas(&circulo)
	*/

	persona1 := models.Persona{}
	persona1.Constructor("Juan", 28)
	fmt.Println(persona1)

	fmt.Println(persona1.GetNombre())
	fmt.Println(persona1.GetEdad())

	persona1.SetNombre("David")
	persona1.SetEdad(29)
	fmt.Println(persona1)

	hellomod.SayHello()
}
