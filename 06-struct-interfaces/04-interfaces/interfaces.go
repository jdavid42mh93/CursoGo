package main

import "fmt"

type Animal interface {
	mover() string
}

type Perro struct {
}

type Pez struct {
}

type Pajaro struct {
}

func moverAnimal(animal Animal) {
	fmt.Println(animal.mover())
}

func (*Perro) mover() string {
	return "soy un perro y camino"
}

func (*Pez) mover() string {
	return "soy un pez y nado"
}

func (*Pajaro) mover() string {
	return "soy un pajaro y vuelo"
}

func main() {
	perro := Perro{}
	moverAnimal(&perro)

	pez := Pez{}
	moverAnimal(&pez)

	pajaro := Pajaro{}
	moverAnimal(&pajaro)
}
