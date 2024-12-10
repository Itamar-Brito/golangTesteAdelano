package presentation

import (
	"github.com/gin-gonic/gin"
)

func NewClientController() *ClientController {
	return &ClientController{}
}

type ClientController struct {
}

func (controller *ClientController) CreateClient(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Client created"})
}
