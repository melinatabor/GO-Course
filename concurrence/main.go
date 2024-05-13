package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// CANALES: Para evitar usar el sleep y esperar a que las goroutines terminen, podemos usar canales.
	// canal := make(chan int) // Declara un canal de tipo entero.
	// canal <- 15             // Envía datos al canal.
	// valor := <-canal        // Recibe datos del canal.

	start := time.Now()

	apis := []string{
		"https://pokeapi.co/",
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}

	// GO ROUTINES: Nos permiten ejecutar funciones de manera concurrente.
	// Sin concurrencia la app tarda de 2 a 3 segundos en ejecutarse.
	// Con concurrencia la app tarda 1 segundo en ejecutarse.
	ch := make(chan string)
	for _, api := range apis {
		// checkAPI(api) // Sin go se ejecuta de manera secuencial.
		go checkAPI(api, ch) // Con go se ejecuta en concurrencia.
	}
	// Imprimimos los mensajes de los canales.
	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}

	// Tiempo de suspenso para que las goroutines terminen y poder ver los resultados.
	// No es la mejor forma porque no sabemos cuánto tiempo tomará la ejecución de las goroutines y podemos terminarlo antes que estas terminen.
	// time.Sleep(1 * time.Second)
	elapsed := time.Since(start)
	fmt.Printf("¡Listo! Tomo %v segundos \n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		// fmt.Printf("ERROR: %s está caído \n", api)
		// Usamos el canal para enviar el mensaje de error:
		ch <- fmt.Sprintf("ERROR: %s está caído \n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s está en funcionamiento \n", api)
}
