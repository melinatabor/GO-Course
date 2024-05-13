package animal

import "fmt"

// INTERFACES: Define un contrato que especifica qué métodos deben ser implementados por otros tipos de datos
type Animal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + "Guau guau")
}

type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + "Miau miau")
}

// Esta funcion nos permite obtener el sonido de cualquier animal que implemente la interfaz Animal
func ObtenerSonido(a Animal) {
	a.Sonido()
}
