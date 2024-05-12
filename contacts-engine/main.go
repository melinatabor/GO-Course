package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Estructura de contactos
// Se puede usar tags para cambiar el nombre de los campos en la serialización y deserialización de JSON
type Contacto struct {
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Telefono string `json:"telefono"`
}

func guardarContacto(contactos []Contacto) error {
	file, err := os.Create("contactos.json")
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file) // Se puede usar json.NewEncoder para serializar un objeto a JSON
	err = encoder.Encode(contactos)  // Se puede usar Encode para serializar un objeto a JSON y escribirlo en un archivo
	if err != nil {
		return err
	}

	return nil
}

// Funcion para obtener los contactos
func obtenerContactos(contactos *[]Contacto) error {
	file, err := os.Open("contactos.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file) // Se puede usar json.NewDecoder para deserializar un objeto desde JSON
	err = decoder.Decode(&contactos) // Se puede usar Decode para deserializar un objeto desde JSON y escribirlo en un archivo
	if err != nil {
		return err
	}

	return nil
}

func main() {

	// Slice de contactos
	var contactos []Contacto

	err := obtenerContactos(&contactos)
	if err != nil {
		fmt.Println("Error al obtener contactos:", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("==== GESTOR DE CONTACTOS ===== \n",
			"1. Agregar contacto\n",
			"2. Mostrar contactos\n",
			"3. Salir\n",
			"Elegi una opción: ")

		var opcion int
		_, err := fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Error al leer la opción:", err)
			return
		}

		switch opcion {
		case 1:
			var c Contacto
			fmt.Print("Nombre: ")
			c.Nombre, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Telefono, _ = reader.ReadString('\n')

			contactos = append(contactos, c)

			if err := guardarContacto(contactos); err != nil {
				fmt.Println("Error al guardar contacto:", err)
			}
		case 2:
			fmt.Println("=========================================================")
			for i, c := range contactos {
				fmt.Printf("%d. %s - %s - %s\n", i+1, c.Nombre, c.Email, c.Telefono)
			}
			fmt.Println("=========================================================")
		case 3:
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}
