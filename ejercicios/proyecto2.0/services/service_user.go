package services

import (
	"fmt"

	userclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/clients"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/dto"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/model"
	errors "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/utils/errors"
)

type userservice struct{}

type userserviceInterface interface {
	GetUserById(id int) (dto.UserDto, errors.ApiError)
	//GetUserByUsername(username string) (dto.UserDto, errors.ApiError)
	GetUsers() (dto.UsersDto, errors.ApiError)
	//InsertUser(dto.UserDto) (dto.UserDto, errors.ApiError)
	Login(cred dto.Credentials) (dto.Token, errors.ApiError)
}

var (
	UserService userserviceInterface
)

func init() {
	UserService = &userservice{}
}

func (s *userservice) GetUserById(id int) (dto.UserDto, errors.ApiError) {

	var usuario model.User = userclient.GetUserById(id)
	var UserDto dto.UserDto

	if usuario.IDuser == 0 {
		return UserDto, errors.NewBadRequestApiError("Bad request")
	}

	UserDto.IDuser = usuario.IDuser
	UserDto.Username = usuario.Username
	UserDto.Firstname = usuario.Firstname
	UserDto.Lastname = usuario.Lastname
	UserDto.Email = usuario.Email
	UserDto.Password = usuario.Password
	UserDto.Phone = usuario.Phone

	return UserDto, nil
}

/*
func (s *userservice) GetUserByUsername(username string) (dto.UserDto, errors.ApiError) {

	var usuario model.User = userclient.GetUserByUsername(username)
	var UserDto dto.UserDto

	if usuario.Username == "" {
		return UserDto, errors.NewBadRequestApiError("Bad request")
	}

	UserDto.IDuser = usuario.IDuser
	UserDto.Username = usuario.Username
	UserDto.Firstname = usuario.Firstname
	UserDto.Lastname = usuario.Lastname
	UserDto.Email = usuario.Email
	UserDto.Password = usuario.Password
	UserDto.Phone = usuario.Phone

	return UserDto, nil
}
*/

func (s *userservice) GetUsers() (dto.UsersDto, errors.ApiError) {

	var usuarios model.Users = userclient.GetUsers()
	var UsersDto dto.UsersDto

	for _, usuario := range usuarios {
		var UserDto dto.UserDto

		UserDto.IDuser = usuario.IDuser
		UserDto.Username = usuario.Username
		UserDto.Firstname = usuario.Firstname
		UserDto.Lastname = usuario.Lastname
		UserDto.Email = usuario.Email
		UserDto.Password = usuario.Password
		UserDto.Phone = usuario.Phone

		UsersDto = append(UsersDto, UserDto)
	}

	return UsersDto, nil
}

/*
func (s *userservice) InsertUser(UserDto dto.UserDto) (dto.UserDto, errors.ApiError) {

	var usuario model.User

	usuario.Username = UserDto.Username
	usuario.Firstname = UserDto.Firstname
	usuario.Lastname = UserDto.Lastname
	usuario.Password = UserDto.Password //hash.GetMD5Hash(UserDto.Password)
	usuario.Phone = UserDto.Phone

	usuario = userclient.InsertUser(usuario)

	UserDto.IDuser = usuario.IDuser

	return UserDto, nil
}
*/
func (s *userservice) Login(cred dto.Credentials) (dto.Token, errors.ApiError) {
	var token dto.Token

	user := userclient.GetUserByUsername(cred)

	if user.IDuser == 0 {
		return token, errors.NewBadRequestApiError("user not found")
	}

	if user.Password != cred.Password {
		return token, errors.NewBadRequestApiError("incorrect password")
	}

	token.Token = fmt.Sprintf("%d", user.IDuser)

	return token, nil
}
