package services

import (
	clients "proyectoarqsoft/clients/users"
	"proyectoarqsoft/dtos"
	"proyectoarqsoft/models"
	encryption "proyectoarqsoft/utils/Encryption"
	"proyectoarqsoft/utils/errors"

	log "github.com/sirupsen/logrus"
)

type userService struct {
	userClient clients.UserClientInterface
}

type userServiceInterface interface {
	Insert(dtos.Userdto) errors.ApiError
	GetUserById(id int) (dtos.Userdto, errors.ApiError)
	GetUserByUsername(username string) (dtos.Userdto, errors.ApiError)
	Login(dtos.Credentials) (dtos.Loginesponsedto, errors.ApiError)
}

var (
	UserService userServiceInterface
)

func initUserService(userClient clients.UserClientInterface) userServiceInterface {
	service := new(userService)
	service.userClient = userClient
	return service
}

func init() {
	UserService = initUserService(clients.UserClient)
}

func (u *userService) GetUserByUsername(username string) (dtos.Userdto, errors.ApiError) {
	var userdto dtos.Userdto

	user, err := u.userClient.GetUserByUsername(username)

	if err != nil {
		return dtos.Userdto{}, errors.NewBadRequestApiError("usuario no encontrado")
	}

	userdto.Id = user.Id
	userdto.Username = user.Username
	userdto.Firstname = user.Firstname
	userdto.Lastname = user.Lastname
	userdto.Password = user.Password

	return userdto, nil

}

func (u *userService) GetUserById(id int) (dtos.Userdto, errors.ApiError) {
	var userdto dtos.Userdto
	user := u.userClient.GetUserById(id)

	if user.Id == 0 {
		return dtos.Userdto{}, errors.NewBadRequestApiError("usuario no valido")
	}

	userdto.Id = user.Id
	userdto.Username = user.Username
	userdto.Firstname = user.Firstname
	userdto.Lastname = user.Lastname
	userdto.Password = user.Password

	return userdto, nil
}

func (u *userService) Insert(userdto dtos.Userdto) errors.ApiError {
	var user models.User

	user.Username = userdto.Username
	user.Firstname = userdto.Firstname
	user.Lastname = userdto.Lastname
	user.Password = encryption.GetSha256Hash(userdto.Password)

	u.userClient.Insert(user)

	return nil
}

func (s *userService) Login(loginDto dtos.Credentials) (dtos.Loginesponsedto, errors.ApiError) {

	hashedpass := encryption.GetSha256Hash(loginDto.Password)

	var user models.User
	user, err := s.userClient.GetUserByUsername(loginDto.Username)
	var loginResponseDto dtos.Loginesponsedto
	loginResponseDto.Id = -1
	if err != nil {
		return loginResponseDto, errors.NewBadRequestApiError("Usuario no encontrado")
	}
	if user.Password != hashedpass && loginDto.Username != "encrypted" {
		return loginResponseDto, errors.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	tokenstring, err := encryption.Signedlogintoken(loginDto)

	if err != nil {
		return loginResponseDto, errors.NewInternalServerApiError("no se pudo obtener el token", err)
	}

	if user.Password != tokenstring && loginDto.Username == "encrypted" {
		return loginResponseDto, errors.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	loginResponseDto.Id = user.Id
	loginResponseDto.Token = tokenstring
	log.Debug(loginResponseDto)
	return loginResponseDto, nil
}
