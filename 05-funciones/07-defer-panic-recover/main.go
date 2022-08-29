package main

import (
	"fmt"
	"os"
)

func main() {

	//funcion
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Oops, al parece el programa no finalizo de forma correcta")
		}
	}()

	if file, error := os.Open("hoa.txt"); error != nil {
		panic("No fue posible obtener el archivo")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo")
			file.Close()
		}()
		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[0:long])
		fmt.Println(contenidoArchivo)
	}

}
