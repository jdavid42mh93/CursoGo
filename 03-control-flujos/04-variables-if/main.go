package main

import "fmt"

func main() {
	if nombre, edad := "Juan", 28; nombre == "Juan" {
		fmt.Printf("Hola %s", nombre)
	} else {
		fmt.Printf("Nombre %s - Edad: %d \n", nombre, edad)
	}

	//obtener valor del map
	users := make(map[string]string)

	users["Juan"] = "juan@gmail.com"
	users["David"] = "david@gmail.com"

	fmt.Println(users)

	if correo, ok := users["Evelyn"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}
}
