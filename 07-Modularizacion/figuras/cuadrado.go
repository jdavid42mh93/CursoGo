package figuras

type Cuadrado struct {
	Lado float32
}

func (cuadrado *Cuadrado) area() float32 {
	return cuadrado.Lado * cuadrado.Lado
}

func (cuadrado *Cuadrado) perimetro() float32 {
	return 4 * cuadrado.Lado
}
