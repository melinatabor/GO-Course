package main

import (
	"fmt"
	"math"
	"strconv"

	"rsc.io/quote"
)

// Declaración de constantes:
const Pi = 3.14
const (
	x = 100
	y = 0b1010 // Binario
	z = 0o12   // Octal
	w = 0xFF   // Hexadecimal
)
const (
	Domingo = iota + 1 // --> Se puede usar iota para asignar valores secuenciales a las constantes
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Hello()) // --> Uso paquete externo

	// Valor predeterminado de las variables: 0 para los números, false para los booleanos y "" para las cadenas

	// Declaración de variables:
	// var firstName, lastName string
	// var age int
	var (
		firstName string = "Melina"
		lastName  string = "Tabor"
		age       int    = 25
	)

	// Forma mas simplificada de declarar variables:
	// var name, lastname, ages = "Melina", "Tabor", 25
	name, lastname, ages := "Melina", "Tabor", 25 // --> Uso de ":=" para evitar usar var, sólo dentro de las funciones se puede

	fmt.Println(firstName, lastName, age)
	fmt.Println(name, lastname, ages)
	fmt.Println(Pi)
	fmt.Println(x, y, z, w)
	fmt.Println(Martes) // --> Se imprime el valor 3

	// Tipo de datos:
	var integer int8 = 127   // --> Entero de 8 bits (el máximo valor que puede tener es 127)
	var uinteger uint = 1    // --> Solo almacena numeros enteros positivos (no tiene signo)
	var float float32 = 3.14 // --> Número decimal de 32 bits, vendria despues del int64
	var bytes byte = 'a'     // --> Enteros sin signo de 8 bits (desde 0 y el máximo valor que puede tener es 255), se usa para caracteres ASCII, los imprime
	var emoji rune = '♡'     // --> Enteros sin signo de 32 bits (desde 0 y el máximo valor que puede tener es 4294967295), se usa para caracteres Unicode, los imprime

	fmt.Println(integer, uinteger, float, bytes, emoji)
	fmt.Println(math.MinInt64, math.MaxInt64) // --> Imprime el valor minimo y maximo de un entero de 64 bits a través del paquete math

	fullName := "Melina Tabor  \"tomar comillas\"" // --> Se puede declarar una variable sin especificar el tipo de dato y Go lo infiere

	fmt.Println(fullName)

	// Conversion de tipos de datos:
	var integer16 int16 = 50
	var integer32 int32 = 100
	fmt.Println(int32(integer16) + integer32) // --> Se convierte el valor de integer16 a int32 para poder sumarlo con integer32

	// Conversiones más complejas: usar el paquete strconv
	s := "100"
	i, _ := strconv.Atoi(s) // --> Convierte un string a un entero, el segundo valor que devuelve es un error, si no se quiere usar se puede usar "_"
	fmt.Println(i + 200)

	n := 42
	s = strconv.Itoa(n) // --> Convierte un entero a un string
	fmt.Println(s + "3")

	// Paquete FMT - Imprimir datos:
	fmt.Print("Hola Mundo\n")                                // --> No agrega un salto de linea
	fmt.Println("Hola Mundo")                                // --> Agrega un salto de linea
	fmt.Printf("Hola soy %s, y tengo %d años\n", name, ages) // --> Se puede usar %s para cadenas y %d para enteros

	greeting := fmt.Sprintf("Hola soy %s, y tengo %d años\n", name, ages) // --> Se puede guardar el string formateado en una variable con Sprintf
	fmt.Println(greeting)

	var nombre string
	var edad int
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanln(&nombre) // --> Se puede usar Scanln para leer un solo string, no mas de uno
	fmt.Print("Ingresa tu edad: ")
	fmt.Scanln(&edad) // --> Se puede usar Scanln para leer un solo entero, no mas de uno
	fmt.Printf("Hola soy %s, y tengo %d años\n", nombre, edad)
	fmt.Printf("Hola soy %v, y tengo %v años\n", nombre, edad) // --> Se puede usar %v para cualquier tipo de dato, solo cuando no sabemos que tipo de dato se va a mostrar
	fmt.Printf(("El tipo de name es: %T \n"), nombre)

	// Operadores aritméticas:
	a := 10
	b := 3
	b++    // --> Incrementa en 1 el valor de b
	b--    // --> Decrementa en 1 el valor de b
	a += 5 // --> Incrementa en 5 el valor de a, es como hacer: a = a + 5
	fmt.Println("Operaciones básicas:", a, b, a+b, a-b, a/b, a*b)
	fmt.Println(math.Pi, math.E) // --> Imprime el valor de Pi y de Euler del paquete math
	fmt.Println(math.Pow(2, 3))  // --> Eleva 2 a la 3, una forma de hacer potencias
	fmt.Println(math.Sqrt(16))   // --> Raiz cuadrada de 16
	fmt.Println(math.Cbrt(27))   // --> Raiz cubica de 27

	// Ejercicio 1: Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y
	// 				luego calcule e imprima el área y perímetro del triángulo.
	fmt.Println("Ejercicio 1:")
	var base, altura, hipotenusa float64
	fmt.Print("Ingresa la base del triángulo: ")
	fmt.Scanln(&base)
	fmt.Print("Ingresa la altura del triángulo: ")
	fmt.Scanln(&altura)
	hipotenusa = math.Sqrt(math.Pow(base, 2) + math.Pow(altura, 2))
	area := (base * altura) / 2
	perimetro := base + altura + hipotenusa

	fmt.Printf("El área del triángulo es: %.2f y el perímetro es: %.2f \n", area, perimetro) // --> Se puede usar %.2f para redondear a 2 decimales

	// Operadores de comparación:
	// Comparación de números
	fmt.Println(1 == 1) // true
	fmt.Println(1 != 2) // true
	fmt.Println(2 < 3)  // true
	fmt.Println(3 > 4)  // false
	fmt.Println(4 <= 4) // true
	fmt.Println(5 >= 6) // false

	// Comparación de cadenas
	fmt.Println("hola" == "hola")  // true
	fmt.Println("hola" != "adios") // true
	fmt.Println("abc" < "def")     // true
	fmt.Println("ghi" > "fgh")     // true
	fmt.Println("hij" <= "hij")    // true
	fmt.Println("klm" >= "klmno")  // false

	// Comparación de booleanos
	fmt.Println(true == true)           // true
	fmt.Println(false != true)          // true
	fmt.Println(true && false == false) // true
	fmt.Println(true || false == true)  // true

	// Operadores de lógicos:
	x := true
	y := false

	// Negación
	fmt.Println(!x) // false
	fmt.Println(!y) // true

	// AND lógico
	fmt.Println(x && x) // true
	fmt.Println(x && y) // false
	fmt.Println(y && y) // false

	// OR lógico
	fmt.Println(x || x) // true
	fmt.Println(x || y) // true
	fmt.Println(y || y) // false

	m := 5
	e := 10
	z := 15
	// Expresión con paréntesis, operadores aritméticos, de comparación y lógicos
	resultado := ((m+e)*z)/(m*e) > z && m != e
	fmt.Println(resultado) //False

	// Declaración If - Else:

}
