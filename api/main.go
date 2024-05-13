package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cd struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Year   int     `json:"year"`
	Price  float64 `json:"price"`
}

// Lo hacemos con datos guardados en memoria, pero se hace con una BD.
var cds = []cd{
	{ID: "1", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Year: 1973, Price: 19.99},
	{ID: "2", Title: "Back in Black", Artist: "AC/DC", Year: 1980, Price: 17.99},
	{ID: "3", Title: "Thriller", Artist: "Michael Jackson", Year: 1982, Price: 18.99},
	{ID: "4", Title: "The Wall", Artist: "Pink Floyd", Year: 1979, Price: 21.99},
	{ID: "5", Title: "Bat Out of Hell", Artist: "Meat Loaf", Year: 1977, Price: 15.99},
}

func main() {
	fmt.Println("¡Bienvenido a la API de CDs!")

	// Creo una nueva instancia de gin que trae configuraciones por defecto, con ella puedo crear rutas y hacer GET o POST.
	router := gin.Default()

	// Al GET se le envia un endpoint y un handler que es una funcion que se ejecuta cuando se hace una peticion a ese endpoint.
	// Como handler le envio una funcion anonima que recibe un contexto de gin y no retorna nada.
	// router.GET("/", func(c *gin.Context) {
	// 	// Le envio un status 200 y un mensaje en formato JSON.
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "¡Hola!",
	// 	})
	// })

	router.GET("/cds", getCDs)
	router.GET("/cds/:id", getCDById)
	router.POST("/cds", postCD)
	router.PUT("/cds/:id", putCD)
	router.DELETE("/cds/:id", deleteCD)
	// Levanto el servidor en el puerto 8080.
	router.Run("localhost:8080")
}

func getCDs(c *gin.Context) {
	// La funcion IndentedJSON recibe un status y un objeto que se serializa a JSON y lo envia en la respuesta.
	c.IndentedJSON(http.StatusOK, cds)
}

func postCD(c *gin.Context) {
	var newCD cd
	if err := c.BindJSON(&newCD); err != nil {
		return
	}
	cds = append(cds, newCD)
	c.IndentedJSON(http.StatusCreated, newCD)
}

func getCDById(c *gin.Context) {
	id := c.Param("id")
	for _, cd := range cds {
		if cd.ID == id {
			c.IndentedJSON(http.StatusOK, cd)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "CD not found"})
}

func deleteCD(c *gin.Context) {
	id := c.Param("id")
	for i, cd := range cds {
		if cd.ID == id {
			cds = append(cds[:i], cds[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, gin.H{"message": "CD deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "CD not found"})
}

func putCD(c *gin.Context) {
	id := c.Param("id")
	var newCD cd
	if err := c.BindJSON(&newCD); err != nil {
		return
	}
	for i, cd := range cds {
		if cd.ID == id {
			cds[i] = newCD
			c.IndentedJSON(http.StatusOK, newCD)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "CD not found"})
}
