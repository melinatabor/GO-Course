package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado en divide:", r)
		}
	}()
	validateZero(divisor)

	return dividendo / divisor, nil
}

// Ejemplo de usar Panic y Recover:
// Se puede usar Panic para lanzar un error y detener la ejecución y Recover para recuperar el error y seguir con la ejecución
// Se puede usar para manejar errores en funciones que no se pueden controlar, que son excepcionales
func validateZero(divisor int) {
	if divisor == 0 {
		// return 0, errors.New("No se puede dividir por cero.") // --> Se puede usar errors.New para crear un nuevo error
		// return 0, fmt.Errorf("No se puede dividir por cero. Dividendo: %d, Divisor: %d", dividendo, divisor) // --> Se puede usar fmt.Errorf para crear un nuevo error con formato
		panic("No se puede dividir por cero.") // --> Se puede usar panic para lanzar un error y detener la ejecución
	}

}

func main() {
	// Manejo de errores:
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(num)

	result, err := divide(10, 0)
	fmt.Println(result, err)

	// Uso de Defer:
	// Se puede usar defer para posponer la ejecución de una función hasta que la función actual retorne.
	// Actua como una PILA de ejecución y el resultado termina siendo 1	, 2 , 3
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close() // Uso defer para cerrar el archivo al finalizar la función y no repetir el close en cada error

	_, err = file.Write([]byte("Hola, soy Meli"))
	if err != nil {
		fmt.Println(err)
		// file.Close()
		return
	}
	// file.Close()

	// Registro de errores:
	// Se puede usar log para registrar errores en un archivo o en la consola
	log.SetPrefix("ERROR: ") // --> Se puede usar SetPrefix para agregar un prefijo a los mensajes de error
	log.Print(err)           // --> log.Print registra el error
	log.Fatal(err)           // --> log.Fatal registra el error y termina la ejecución del programa
	log.Panic(err)           // --> log.Panic registra el error y termina la ejecución del programa

	logger, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) // --> Se puede usar OpenFile para abrir un archivo y escribir en él
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logger) // --> Se puede usar SetOutput para escribir los logs en un archivo
	log.Print("Soy un log")

	defer logger.Close()
}
