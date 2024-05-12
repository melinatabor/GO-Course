package main

import (
	"fmt"
	"log"

	"github.com/melinatabor/GO-Course/modules"
)

// Para crear un ejecutable de este archivo uso `go build main.go` y luego `./main`
// Si le quiero cambiar el nombre al ejecutable uso `go build -o hello main.go` y luego `./hello`
func main() {
	log.SetPrefix("Hello: ") // SetPrefix se usa para cambiar el prefijo de los mensajes de registro
	log.SetFlags(0)          // SetFlags se usa para cambiar la forma en que se muestran los mensajes de registro

	names := []string{"Meli", "Tabor", "Golang"}
	messages, err := modules.Hellos(names)
	// message, err := modules.Hello("Meli")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
