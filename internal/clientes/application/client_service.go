package application

import (
	"main/internal/clientes/domain"
)

type ClientService struct {
}

func NewClientService() *ClientService {
	return &ClientService{}
}

func (service *ClientService) CreateClient(clientToBeCreated domain.Client) (domain.Client, error) {
	//cliação do cliente no banco

	return clientToBeCreated, nil
}
