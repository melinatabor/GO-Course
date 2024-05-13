package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	PrintList("Meli", "Hola", "Mundo", "GO", "Course", 3.14, true, 5, 10, 15, 20, 25, 30, 35, 100)
	PrintList2("Meli", "Hola", "Mundo", "GO", "Course", 3.14, true, 5, 10, 15, 20, 25, 30, 35, 100)

	fmt.Println(Suma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(SumaConRestriccion[float32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(SumaConRestriccionDeGo[uint](1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Includes(strings, "a"))
	fmt.Println(Includes(numbers, 6))

	fmt.Println(Filter(numbers, func(n int) bool {
		return n > 3
	}))
	fmt.Println(Filter(strings, func(s string) bool {
		return s == "a"
	}))

	producto := Product[uint]{
		Id:          1,
		Description: "Producto 1",
		Price:       100.00,
	}
	producto2 := Product[string]{
		Id:          "2",
		Description: "Producto 2",
		Price:       200.00,
	}
	fmt.Println(producto, producto2)

	fmt.Println(Sum(1, 2, 3, 4.5, 5.5, 6))

	var nun1 integere = 10
	var nun2 integere = 20
	fmt.Println(Sum2(nun1, nun2))
}

// MISMA MANERA DE HACER LO MISMO CON ANY Y INTERFACES
// INTERFACES: Son un tipo de dato que puede ser cualquier cosa.
func PrintList(list ...interface{}) {
	for _, value := range list {
		fmt.Println(value)
	}
}

// ANY: Es un tipo de dato que puede ser cualquier cosa.
func PrintList2(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

// RESTRICCIONES: Son para restringir los tipos de datos que se pueden usar en una funcion.
// Para que no nos queden todos los tipos de datos ahi, podemos hacer una restriccion.
func Suma[T ~int | ~float64 | ~float32 | ~uint](nums ...T) T { // Significa que T puede ser int o float64 usando numeros negativos o positivos.
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Restriccion con interfaces
type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

func SumaConRestriccion[T Numbers](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Reutilizamos la restriccion de GO del paquete
func SumaConRestriccionDeGo[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// RESTRICCIONES CON OPERADORES: Se pueden usar operadores como == o != para restringir los tipos de datos.
// El T Comparable es una restriccion de GO que se puede usar para comparar tipos de datos.
func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// Funcion que recibe una lista de elementos y una funcion que retorna un booleano.
func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}

// ESTRUCTURAS GENÉRICAS: Se pueden usar estructuras genéricas para reutilizar código, y si tenemos una bd relacional que autoincrementa el id, podemos usar un tipo de dato genérico.
// Pero si tenemos una bd documental que no autoincrementa el id, podemos usar un tipo de dato string.
// Representa la tabla en la base de datos y representa el objeto de un producto.
type Product[T uint | string] struct {
	Id          T
	Description string
	Price       float64
}

// FUNCIONES GENÉRICAS:
// Parámetros de tipos son los que se usan en la definición de la función.
func Sum[T int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// RESTRICCIÓN DE APROXIMACIÓN: Se puede usar para restringir los tipos de datos que se pueden usar en una función.
type integere int

// Usando el ~ se puede hacer una aproximación de los tipos de datos y no tener que especificar los tipos de datos nuevos que cree si derivan de esos.
// En este caso no es necesario agregar integer porque deriva de int.
func Sum2[T ~int | ~float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}
