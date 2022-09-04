package main

import "fmt"

type Figura interface {
	area() float32
	perimetro() float32
}

type Cuadrado struct {
	lado float32
}

type Circulo struct {
	radio float32
}

func medidas(figura Figura) {
	fmt.Println(figura.area())
	fmt.Println(figura.perimetro())
}

func (cuadrado *Cuadrado) area() float32 {
	return cuadrado.lado * cuadrado.lado
}

func (cuadrado *Cuadrado) perimetro() float32 {
	return 4 * cuadrado.lado
}

func (circulo *Circulo) area() float32 {
	return 3.14 * circulo.radio * circulo.radio
}

func (circulo *Circulo) perimetro() float32 {
	return 2 * 3.14 * circulo.radio
}

func main() {
	cuadrado := Cuadrado{lado: 5}
	medidas(&cuadrado)

	circulo := Circulo{radio: 3}
	medidas(&circulo)
}
