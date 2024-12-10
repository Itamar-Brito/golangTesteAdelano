package main

import (
	"main/internal/clientes/presentation"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	clientController := presentation.NewClientController()

	r.GET("/client", clientController.CreateClient)

	r.Run(":8083")
}
