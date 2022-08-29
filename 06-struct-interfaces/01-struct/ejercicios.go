package main

import "fmt"

type Pais struct {
	nombre      string
	cantidadhab int
}

func main() {
	pais1 := Pais{"Ecuador", 30}
	pais2 := Pais{"Colombia", 50}
	pais3 := Pais{"Brasil", 100}
	if pais1.cantidadhab > pais2.cantidadhab && pais1.cantidadhab > pais3.cantidadhab {
		fmt.Printf("Pais con mayor cantidad de habitantes: %s \n", pais1.nombre)
	} else {
		if pais2.cantidadhab > pais3.cantidadhab {
			fmt.Printf("Pais con mayor cantidad de habitantes: %s\n", pais2.nombre)
		} else {
			fmt.Printf("Pais con mayor cantidad de habitantes: %s \n", pais3.nombre)
		}
	}
}
