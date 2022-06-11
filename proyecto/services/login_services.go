package services

import (
	//"strings"

	clientclient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/client"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"

	//errors "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/utils/errors"
	"errors"

	//"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/utils/files"
	hash "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/utils/hash"
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

//login trucho
func login(cred dto.Credentials) (dto.Hash, error) {
	var client model.Client
	client = clientclient.GetClientByUsername(cred.Client)

	loggedIn := false
	hash, _ := hash.HashPassword(cred.Password)

	if cred.Client == client.Username && hash == client.Hash {
		loggedIn = true
	}

	if !loggedIn {
		return dto.Hash{}, errors.New("credenciales no validas")
	}
	return dto.Hash{
		Hash: hash,
	}, nil

}
