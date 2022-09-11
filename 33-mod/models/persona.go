package models

type Persona struct {
	name string
	edad int
}

func (p *Persona) Constructor(name string, edad int) {
	p.name = name
	p.edad = edad
}
