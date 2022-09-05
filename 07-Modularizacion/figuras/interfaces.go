package figuras

import "fmt"

type figura interface {
	area() float32
	perimetro() float32
}

func Medidas(figura figura) {
	fmt.Println(figura.area())
	fmt.Println(figura.perimetro())
}
