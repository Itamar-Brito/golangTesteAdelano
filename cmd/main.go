package main

import (
	presentationclient "main/internal/clientes/presentation"
	presentationveiculos "main/internal/veiculos/presentation"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	clientRoutes := r.Group("/client")
	{
		clientController := presentationclient.NewClientController()
		clientRoutes.POST("/create", clientController.CreateClient)
	}

	veiculosRoutes := r.Group("/veiculos")
	{
		veiculosController := presentationveiculos.NewVeiculosController()
		veiculosRoutes.POST("/create", veiculosController.CreateVeiculos)
	}

	r.Run(":8083")
}
