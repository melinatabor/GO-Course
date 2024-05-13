package main

import "fmt"

func main() {
	suma("Meli", 1, 45, 30, 10)
	imprimirDatos("Meli", 1, 45, 30, 10, "Hola", 3.14, true)
	fmt.Println(factorial(5))

	// FUNCIONES ANÓNIMAS: Son funciones sin nombre que se asignan a una variable.
	saludo := func(name string) {
		fmt.Printf("Hola %s desde una función anónima \n", name)
	}

	saludar("Meli", saludo)
	r1 := aplicar(duplicar, 5) // Pasamos la función como parámetro.
	r2 := aplicar(triplicar, 5)
	fmt.Println(r1, r2)

	r := double(addOne, 2)
	fmt.Println("Resultado", r)

	// CLOSURES: Va a recordar el valor de x y lo va a incrementar cada vez que se llame a la función.
	nextInt := incrementar()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

// FUNCIONES VARIÁDICAS: Reciben n parámetros del mismo tipo con la ayuda de los puntos suspensivos.
func suma(nombre string, numeros ...int) int {
	// fmt.Printf("%T - %v \n", numeros, numeros)

	suma := 0
	for _, numero := range numeros {
		suma += numero
	}
	fmt.Printf("Hola %s, el resultado de la suma es: %d \n", nombre, suma)
	return suma
}

// Funcion variádica con interface{} para recibir diferentes tipos de datos de n cantidades.
func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Println(dato)
	}
}

// FUNCIONES RECURSIVAS: Son funciones que se llaman a sí mismas generando un bucle.
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// FUNCIONES COMO PARÁMETROS: Podemos pasar funciones como parámetros a otras funciones usando funciones anónimas.
func saludar(name string, f func(string)) {
	f(name)
}

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}

// FUNCIONES DE ORDEN SUPERIOR: Son funciones que reciben otras funciones como parámetros y devuelve una función como resultado.
func double(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int {
	return x + 1
}

// CLOSURES: Son funciones anónimas que se definen dentro de otra función y que pueden acceder a las variables locales de la función que las contiene.
// Esto significa que un closure puede recordar el valor de la variable del ámbito en el que fue creado y utilizarlas en cualquier momento posterior.

// Funcione que retorna una función que retorna un entero.
func incrementar() func() int {
	x := 0
	// Retorna una función anónima que incrementa la variable x.
	return func() int {
		x++
		return x
	}
}
