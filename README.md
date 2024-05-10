# GO
## Caracteristicas:
- Código abierto
- Compilado
- Tipado Estático
- Multiplataforma y multiparadigma
- Manejo de concurrencia: Permite que cualquier función se ejecute como un subproceso ligero, simplemente usando la palabra reservada go para implementar la concurrencia
- Biblioteca estándar muy grande
- Simple, sencillo y fácil de aprender
- Administración automática de la memoria, como la recoleccion de elementos no utilizados
- Hereda aspectos de C
- Comparte similitudes con Java y Python
- Algunas caracteríticas de POO pero simplificadas

Documentación: https://go.dev/doc

## Playground de GO:
No podemos utilizar paquetes externos, si el usuario tiene que interactuar tampoco vamos a poder hacer entradas.
````
// Todos los archivos de GO deben pertenecer a un paquete. Aca indicamos que pertenece al paquete principal
package main

// Importamos el paquete estándar de entrada y salida de datos
import "fmt"

// Aca indicamos que es la funcion principal con la palabra reservada main. Desde esta función es desde donde inicia a ejecutarse una app
func main() {
	fmt.Println("Hola mundo ;), 世界")
}
````

## Comandos
`go run nombre-archivo.go` Nos corre la funcion
`go build nombre-archivo.go` Nos genera un exe y ya podemos correrlo de forma mas sencilla
`./nombre-archivo`

## Paquetes de Terceros (Externos)
Tenemos que inicializar un manejador de módulos para nuestra app: `go mod init nombre-manejador-modulo-que-quiera`
Se crea un archivo `go.mod` que contiene el nombre del manejador de módulos y la versión de go. 
Se utiliza para definir y gestionar los módulos y las dependencias de mi proyecto en go.
Cuando descarguemos archivos externos se genera el archivo `go.sum` que registra la suma de verificaciones de los módulos y dependencias de proyecto.

Instalar paquete externo: `go get rsc.io/quote` y se agrega al go.mod
Luego lo puedo importar:
````
import "rsc.io/quote"

func main() {
    fmt.Println(quote.Hello())
}
````