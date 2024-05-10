package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
	Email  string
}

func main() {
	// ARRAYS: Estructura de datos fija que almacena elementos del mismo tipo
	var array = [5]int{10, 20, 5, 25, 2} // --> Declaración de un array que va a almacenar 5 elementos
	array[0] = 1

	// Se puede usar len(array) para obtener la longitud del array
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
	// Se puede usar range para recorrer el array
	for index, value := range array {
		fmt.Println("Index:", index, "Value:", value)
	}

	// MATRICES: Estructura de datos fija que almacena elementos del mismo tipo en filas y columnas
	var matrix = [3][3]int{{0, 2, 3}, {4, 5, 6}, {7, 8, 9}} // --> Declaración de una matriz de 3x3
	matrix[0][0] = 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}

	// SLICES: Estructura de datos dinámica (se pueden agregar o eliminar elementos) que almacena elementos del mismo tipo
	var slice = []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"} // --> Declaración de un slice que va a almacenar 5 elementos
	dia := slice[0:5]

	dia = append(dia, "Viernes")         // --> Se puede usar append(slice, elemento) para agregar un elemento al final del slice
	dia = append(dia[:2], dia[3:]...)    // --> Se puede usar append(slice[:i], slice[j:]...) para eliminar un elemento del slice
	fmt.Println(dia, len(dia), cap(dia)) // --> Se puede usar len(slice) para obtener la longitud del slice y cap(slice) para obtener la capacidad del slice

	nombres := make([]string, 5, 10) // --> Se puede usar make([]tipo, longitud, capacidad) para crear un slice con una longitud y capacidad específica
	nombres[0] = "Melina"
	fmt.Println(nombres, len(nombres), cap(nombres))

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5)

	copy(slice2, slice1) // --> Se puede usar copy(destino, origen) para copiar los elementos de un slice a otro
	fmt.Println(slice2)

	// MAPAS: Estructura de datos que almacena elementos en pares clave-valor
	colors := map[string]string{
		"rojo":  "#ff0000",
		"verde": "#00ff00",
		"azul":  "#0000ff",
	} // --> Declaración de un mapa que va a almacenar colores y sus códigos hexadecimales

	colors["amarillo"] = "#ffff00" // --> Se puede usar mapa[clave] = valor para agregar un nuevo par clave-valor al mapa

	valorRojo, ok := colors["rojo"] // --> Se puede acceder a un valor específico del mapa usando la clave correspondiente y obtener un booleano que indica si la clave existe
	delete(colors, "azul")          // --> Se puede usar delete(mapa, clave) para eliminar un par clave-valor del mapa
	for clave, valor := range colors {
		fmt.Println(clave, valor)
	} // --> Se puede usar range para recorrer el mapa
	fmt.Println(colors, valorRojo, ok) // --> Se puede acceder a un valor específico del mapa usando la clave correspondiente

	// ESTRUCTURAS: Estructura de datos que almacena elementos de diferentes tipos
	// Una estructura seria el modelo de una clase en otros lenguajes de programación, y realizar instancias de la misma seria crear objetos de esa clase

	// type Persona struct {
	// 	Nombre string
	// 	Edad   int
	// 	Email  string
	// }
	var persona Persona // --> Declaración de una instancia de tipo Persona
	persona.Nombre = "Melina"
	persona.Edad = 25
	persona.Email = "melina.tabor@mercadolibre.com"

	persona2 := Persona{"Alex", 30, "alex@yahoo.com"} // --> Declaración de una instancia de tipo Persona con valores iniciales
	fmt.Println(persona, persona2)

	// MÉTODOS Y PUNTEROS:
	// Un método es una función que tienen un receptor que es un puntero o una variable de la estructura
	// Un puntero es una variable que almacena la dirección de memoria de otra variable. Se usan para referenciar y acceder a la variable original.
	// Los punteros son utilizados como receptores de métodos para modificar el valor de la variable original.

	var x int = 10
	var p *int = &x // --> Declaración de un puntero que almacena la dirección de memoria de la variable x

	editar(&x)            // --> Se puede usar &x para enviar la dirección de memoria de la variable x a la función editar
	fmt.Println(x, p, *p) // --> Se puede usar *p para acceder al valor de la variable a la que apunta el puntero

	persona.hello() // --> Los métodos se deben acceder mediante la instancia de la estructura, no son independientes como las funciones
}
func (p *Persona) hello() { // --> Para que la función sea un método, se debe agregar el receptor (p *Persona) antes del nombre de la función
	// Recibe un puntero de tipo Persona
	fmt.Println("Hola", p.Nombre)
}
func editar(x *int) {
	*x = 20 // --> Se puede usar *x para modificar el valor de la variable a la que apunta el puntero
}
