package main

import "fmt"

// Struct Persona
type Persona struct {
	nombre string
	edad   int
}

type Empleado struct {
	Persona
	sueldo float32
}

// Metodos
func (p *Persona) impirmir() {
	fmt.Printf("Nombre: %s\n Edad: %d\n", p.nombre, p.edad)
}

func (p *Persona) editarNombre(nombre string) {
	p.nombre = nombre
}

func main() {
	p1 := Persona{"Juan", 28}
	p1.impirmir()
	p1.editarNombre("David!")

	p1.impirmir()
	p2 := Persona{
		nombre: "Evelyn",
		edad:   25,
	}
	p2.editarNombre("Andrea")
	p2.impirmir()

	emp1 := Empleado{}
	emp1.nombre = "Segundo"
	emp1.edad = 56
	emp1.sueldo = 500.0
	emp1.impirmir()
	fmt.Println(emp1)
}
