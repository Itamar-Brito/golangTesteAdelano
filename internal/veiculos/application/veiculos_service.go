package application

import "main/internal/veiculos/domain"

type VeiculosService struct {
}

func NewVeiculosService() *VeiculosService {
	return &VeiculosService{}
}

func (service *VeiculosService) CreateVeiculos(VeiculosToBeCreated domain.Veiculos) (domain.Veiculos, error) {
	//cliação do Veiculose no banco

	return VeiculosToBeCreated, nil
}
