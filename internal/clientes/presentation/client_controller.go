package presentation

import (
	"main/internal/clientes/application"
	"main/internal/clientes/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewClientController() *ClientController {
	return &ClientController{}
}

type ClientController struct {
}

func (controller *ClientController) CreateClient(c *gin.Context) {

	var client domain.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	clientService := application.NewClientService()
	clientAddened, err := clientService.CreateClient(client)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Client created successfully", "client": clientAddened})
}
