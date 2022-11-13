package services

import (
	//"strings"

	clientclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/client"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"
	errors "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/utils/errors"
	"github.com/dgrijalva/jwt-go"
	//errors "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/utils/errors"
	//"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/utils/files"
	//hash "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/utils/hash"
)

/*
const (
	credentialsPath = "credentials.txt"
	tokenPath = "token.txt"
)
*/

/*
func Login(cred dto.Credentials)  (dto.Token, error) {
	bytes, err := utils.ReadFile(credentialsPath)
	if err != nil {
		return dto.Token{}, err
	}

	loggedIn := false
	for _, line := range strings.Split(string(bytes), "\n") {
		components := strings.Split(line, "@")
		user, password := components[0], components[1]
		if user == cred.Client && password == cred.Password {
			loggedIn = true
			break
		}
	}

	if !loggedIn {
		return dto.Token{}, errors.New("invalid credentials")
	}

	tokenBytes, err := utils.ReadFile(tokenPath)
	if err != nil {
		return dto.Token{}, err
	}

	return dto.Token{
		Token: string(tokenBytes),
	}, nil
}
*/

/*
//login satanico, preguntar
func Login(cred dto.Credentials) (string, error) {
	var client model.Client
	client = clientclient.GetClientByUsername(cred.Client)

	loggedIn := false
	password := hash.GetMD5Hash(cred.Password)

	if cred.Client == client.Username && password == client.Password {
		loggedIn = true
	}

	if !loggedIn {
		return password, errors.New("credenciales no validas")
	}
	return password, nil
}
*/

var jwtKey = []byte("secret_key")

func Login(cred dto.Credentials) (dto.Token, errors.ApiError) {
	//var user model.User
	var client model.Client = clientclient.GetClientByUsername(cred.Client) //objeto de la DB, a traves del DAO
	//var userDto dto.UserDto
	var tokenDto dto.Token

	if client.ID_client == 0 {
		return tokenDto, errors.NewBadRequestApiError("user not found")
	}

	if client.Password == cred.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
	}
	return tokenDto, nil
}
