package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

type ListaTareas struct {
	tareas []Tarea
}

// MÉTODO PARA AGREGAR TAREA:
// El receptor es un puntero a ListaTareas, esto significa que el método agregarTarea esta asociado con el tipo ListaTareas
// Instancia es la instancia del tipo ListaTareas que se esta utilizando, dentro del cual se puede acceder a los campos y métodos del tipo
func (instancia *ListaTareas) agregarTarea(tarea Tarea) {
	instancia.tareas = append(instancia.tareas, tarea)
}

// MÉTODO PARA EDITAR TAREAS:
func (instancia *ListaTareas) editarTarea(indice int, tarea Tarea) {
	instancia.tareas[indice] = tarea
}

// MÉTODO PARA ELIMINAR TAREAS:
func (instancia *ListaTareas) eliminarTarea(indice int) {
	instancia.tareas = append(instancia.tareas[:indice], instancia.tareas[indice+1:]...)
}

// MÉTODO PARA MARCAR COMO COMPLETAS LAS TAREAS:
func (instancia *ListaTareas) marcarCompletado(indice int) {
	instancia.tareas[indice].completado = true
}

func main() {
	// Instancia de ListaTareas
	lista := ListaTareas{}

	leer := bufio.NewReader(os.Stdin) // --> Se puede usar bufio.NewReader para leer datos del usuario
	for {
		var opcion int
		fmt.Println(
			"Seleccione una opción: \n",
			"1. Agregar tarea \n",
			"2. Editar tarea \n",
			"3. Eliminar tarea \n",
			"4. Marcar como completada \n",
			"5. Salir",
		)
		fmt.Print("Ingrese la opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n') // --> Se puede usar ReadString para leer un string del usuario
			fmt.Print("Ingrese la descripción de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
		case 2:
			var indice int
			var t Tarea
			fmt.Print("Ingrese el índice de la tarea a editar: ")
			fmt.Scanln(&indice)
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripción de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.editarTarea(indice-1, t)
		case 3:
			var indice int
			fmt.Print("Ingrese el índice de la tarea a eliminar: ")
			fmt.Scanln(&indice)
			lista.eliminarTarea(indice - 1)
		case 4:
			var indice int
			fmt.Print("Ingrese el índice de la tarea a marcar como completada: ")
			fmt.Scanln(&indice)
			lista.marcarCompletado(indice - 1)
		case 5:
			fmt.Println("Saliendo...")
			os.Exit(0)
		default:
			fmt.Println("Opción inválida")

		}

		// Listar tareas
		fmt.Println("Lista de tareas:")
		fmt.Println("================================================================")
		for i, tarea := range lista.tareas {
			fmt.Printf("%d. %s %s Completado: %t\n", i+1, tarea.nombre, tarea.descripcion, tarea.completado)
		}
		fmt.Println("================================================================")
	}
}
