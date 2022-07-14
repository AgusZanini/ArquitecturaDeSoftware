package clients

import (
	"fmt"

	//"github.com/AgusZanini/ArquitecturaDeSoftware/mvc-go/model"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/dto"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/model"

	//log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DbUser *gorm.DB

func GetUserById(id int) model.User {
	var usuario model.User
	if DbUser == nil {
		fmt.Println("nil User")
	}
	if result := DbUser.Where("IDuser = ?", id).First(&usuario); result.Error != nil {
		fmt.Println("couldn't find user")
	}
	fmt.Println("User: ", usuario)

	return usuario
}

func GetUserByUsername(cred dto.Credentials) model.User {
	var usuario model.User

	if DbUser == nil {
		fmt.Println("nil user")
	}

	if result := DbUser.Where("Username = ?", cred.Username).First(&usuario); result.Error != nil {
		fmt.Println("couldn't find user by username")
	}

	fmt.Println("User: ", usuario)

	return usuario
}

func GetUsers() model.Users {
	var usuarios model.Users

	if result := DbUser.Find(&usuarios); result.Error != nil {
		fmt.Println("error finding users")
	}
	fmt.Println("usuarios: ", usuarios)

	return usuarios
}

/*
func InsertUser(usuario model.User) model.User {
	result := DbUser.Create(&usuario)

	if result.Error != nil {
		fmt.Println("error")
	}

	fmt.Println("usuario agregado", usuario.IDuser)
	return usuario
}
*/
