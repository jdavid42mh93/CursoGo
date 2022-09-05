package figuras

type Circulo struct {
	Radio float32
}

func (circulo *Circulo) area() float32 {
	return 3.14 * circulo.Radio * circulo.Radio
}

func (circulo *Circulo) perimetro() float32 {
	return 2 * 3.14 * circulo.Radio
}
