package presentation

import (
	"main/internal/veiculos/application"
	"main/internal/veiculos/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewVeiculosController() *VeiculosController {
	return &VeiculosController{}
}

type VeiculosController struct {
}

func (controller *VeiculosController) CreateVeiculos(c *gin.Context) {

	var Veiculos domain.Veiculos

	if err := c.ShouldBindJSON(&Veiculos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	VeiculosService := application.NewVeiculosService()
	VeiculosAddened, err := VeiculosService.CreateVeiculos(Veiculos)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Veiculos created successfully", "Veiculos": VeiculosAddened})
}
