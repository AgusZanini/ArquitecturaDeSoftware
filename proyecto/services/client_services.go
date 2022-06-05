package services

import (
	clientClient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/client"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"
	errors "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/utils/errors"
)

type clientservice struct{}

type clientserviceInterface interface {
	GetClientById(id int) (dto.ClientDto, errors.ApiError)
	GetClients() (dto.ClientsDto, errors.ApiError)
	InsertClient(dto.ClientDto) (dto.ClientDto, errors.ApiError)
}

var (
	ClientService clientserviceInterface
)

func init() {
	ClientService = &clientservice{}
}

func (s *clientservice) GetClientById(id int) (dto.ClientDto, errors.ApiError) {

	var cliente model.Client = clientClient.GetClientById(id)
	var clientDto dto.ClientDto

	if cliente.ID_client == 0 {
		return clientDto, errors.NewBadRequestApiError("Bad request") //como devolver un error
	}

	clientDto.ID_client = cliente.ID_client
	clientDto.First_name = cliente.First_name
	clientDto.Last_name = cliente.Last_name
	clientDto.Email = cliente.Email
	clientDto.Password = cliente.Password
	clientDto.Phone = cliente.Phone

	return clientDto, nil
}

func (s *clientservice) GetClients() (dto.ClientsDto, errors.ApiError) {

	var clientes model.Clients = clientClient.GetClients()
	var clientsDto dto.ClientsDto

	for _, cliente := range clientes {
		var clientDto dto.ClientDto

		clientDto.ID_client = cliente.ID_client
		clientDto.First_name = cliente.First_name
		clientDto.Last_name = cliente.Last_name
		clientDto.Email = cliente.Email
		clientDto.Password = cliente.Password
		clientDto.Phone = cliente.Phone

		clientsDto = append(clientsDto, clientDto) // preguntar
	}

	return clientsDto, nil
}

func (s *clientservice) InsertClient(clientDto dto.ClientDto) (dto.ClientDto, errors.ApiError) {

	var cliente model.Client

	cliente.First_name = clientDto.First_name
	cliente.Last_name = clientDto.Last_name
	cliente.Password = clientDto.Password
	cliente.Phone = clientDto.Phone

	//En el ejemplo de edu hace estas dos lineas, preguntar por que

	cliente = clientClient.InsertClient(cliente)

	clientDto.ID_client = cliente.ID_client

	return clientDto, nil
}
