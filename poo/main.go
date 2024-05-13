package main

import (
	"fmt"
	"poo/animal"
	"poo/book"
)

func main() {
	// Creo una instancia de un libro con sus atributos
	// No se puede acceder a los atributos porque son privados
	// var myBook = book.Book{
	// 	Titulo:  "El principito",
	// 	Autor:   "Antoine de Saint-Exupéry",
	// 	Paginas: 100,
	// }

	// Creo una instancia de un libro con el constructor
	myBook := book.NewBook("El principito", "Antoine de Saint-Exupéry", 100)
	myBook.PrintInfo()

	myBook.SetTitulo("El principito (edicion limitada)")
	fmt.Println(myBook.GetTitulo())

	myTextBook := book.NewTextBook("Matematicas 1", "Juanito Perez", 200, "Santillana", "Primaria")
	myTextBook.PrintInfo()

	// Uso de polimorfismo con la funcion Print
	book.Print(myBook)
	book.Print(myTextBook)

	perro := animal.Perro{Nombre: "Firulais "}
	gato := animal.Gato{Nombre: "Garfield "}
	animal.ObtenerSonido(&perro)
	animal.ObtenerSonido(&gato)

	animales := []animal.Animal{
		&perro,
		&gato,
		&animal.Perro{Nombre: "Hannah "},
		&animal.Gato{Nombre: "Mittens "},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
