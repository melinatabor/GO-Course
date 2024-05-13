package book

import "fmt"

// ENCAPSULACION: Se refiere a la proteccion de los atributos de una clase, para que no puedan ser accedidos directamente desde fuera de la clase
// Objeto Book con sus atributos privados para mantener el encapsulamiento, para ser publicos deben comenzar con mayuscula
type Book struct {
	titulo  string // Atributo privado
	autor   string // Atributo privado
	paginas int    // Atributo privado
}

// Constructor de la clase Book
func NewBook(titulo string, autor string, paginas int) *Book {
	return &Book{
		titulo:  titulo,
		autor:   autor,
		paginas: paginas,
	}
}

// GET and SET: Metodos para acceder a los atributos privados
func (b *Book) SetTitulo(titulo string) {
	b.titulo = titulo
}

func (b *Book) GetTitulo() string {
	return b.titulo
}

// Metodo de la clase Book
func (b *Book) PrintInfo() {
	fmt.Printf("Titulo: %s \n Autor: %s \n Paginas: %d \n", b.titulo, b.autor, b.paginas)
}

// COMPOSICION: Alternativa a la herencia, permite crear nuevas estructuras que contienen campos y metodos de otras estructuras
type TextBook struct {
	Book      // Aca se incluyen los campos y metodos de la clase Book
	editorial string
	niveles   string
}

// Constructor de la clase TextBook que hereda de Book y agrega sus propios atributos
func NewTextBook(titulo string, autor string, paginas int, editorial string, niveles string) *TextBook {
	return &TextBook{
		Book: Book{
			titulo:  titulo,
			autor:   autor,
			paginas: paginas,
		},
		editorial: editorial,
		niveles:   niveles,
	}
}

// Metodo de la clase TextBook que hereda de Book
func (t *TextBook) PrintInfo() {
	fmt.Printf("Titulo: %s \n Autor: %s \n Paginas: %d \n Editorial: %s \n Niveles: %s \n", t.titulo, t.autor, t.paginas, t.editorial, t.niveles)
}

// POLIMORFISMO: Capacidad de una clase de tener metodos con el mismo nombre pero con diferentes parametros
// Interfaz es un conjunto de métodos que una estructura debe implementar para satisfacer la interfaz.
// Esto permite que una estructura implemente multiples interfaces y obtenga diferentes comportamientos.

type Printable interface {
	PrintInfo() // Este interface dice que cada estructura debe tener este método para ser considerada Printable
}

// Funcion que recibe un objeto que implementa la interfaz Printable
func Print(p Printable) {
	p.PrintInfo()
}
