package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	// IF: Se pueden declarar variables dentro de la condición
	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("Buenos días!")
	} else if t.Hour() < 17 {
		fmt.Println("Buenas tardes!")
	} else {
		fmt.Println("Buenas noches!")
	}

	// SWITCH: No es necesario colocar los break, ya que se hace de forma automática
	// runtime.GOOS es una variable que contiene el sistema operativo donde se está ejecutando el programa go
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Ejecutando en Windows")
	case "linux":
		fmt.Println("Ejecutando en Linux")
	case "darwin":
		fmt.Println("Ejecutando en MacOS")
	default:
		fmt.Println("Sistema operativo desconocido")
	}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Buenos días!")
	case t.Hour() < 17:
		fmt.Println("Buenas tardes!")
	default:
		fmt.Println("Buenas noches!")
	}

	// FOR: Es la única estructura repetitiva en Go
	// Bucle infinito
	for {
		fmt.Println("Hola Mundo!")
		break
	}

	// Bucle con condición
	var i int
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Bucle con inicialización, condición y actualización
	for j := 0; j < 10; j++ {
		fmt.Println(j)
		if j == 5 {
			break // --> Se puede usar break para salir del bucle
		}
		if j == 3 {
			continue // --> Se puede usar continue para saltar a la siguiente iteración
		}
	}

	// Llamada a las funciones:
	hello("Meli")
	add(5, 3)
	fmt.Println(calc(2, 2))
	game()
}

// Creación de funciones:
// Función que recibe parámetros
func hello(name string) {
	fmt.Println("Hola " + name)
}

// Función que retorna un valor
func add(a, b int) int {
	fmt.Println(a + b)
	return a + b
}

// Función que retorna más de un valor
func calc(a, b int) (sum, rest int) {
	sum = a + b
	rest = a - b
	return
}

func game() {
	rand.Seed(time.Now().UnixNano()) // --> Se puede usar para generar números aleatorios

	numAleatorio := rand.Intn(100) // --> Genera un número aleatorio entre 0 y 99
	var numIngresado, intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf(("Ingresa un número (%d intentos restantes): "), maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Printf("¡Felicidades! Adivinaste el número en %d intentos\n", intentos)
			playAgain()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El número ingresado es mayor al número aleatorio")
		} else {
			fmt.Println("El número ingresado es menor al número aleatorio")
		}
	}

	fmt.Printf("¡Lo siento! No adivinaste el número en %d intentos. El número era: %d\n", maxIntentos, numAleatorio)
	playAgain()
}

func playAgain() {
	var eleccion string
	fmt.Print("¿Quieres jugar de nuevo? (s/n): ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		game()
	case "n":
		fmt.Println("¡Gracias por jugar!")
	default:
		fmt.Println("Opción no válida")
		playAgain()
	}
}
