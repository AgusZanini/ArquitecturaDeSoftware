package clients

import (
	"fmt"
	"proyectoarqsoft/models"

	"gorm.io/gorm"
)

var Db *gorm.DB

type userClient struct{}

type UserClientInterface interface {
	Insert(user models.User)
	GetUserById(id int) models.User
	GetUserByUsername(username string) (models.User, error)
}

var UserClient UserClientInterface

func init() {
	UserClient = &userClient{}
}

func (u *userClient) GetUserByUsername(username string) (models.User, error) {
	var user models.User

	result := Db.Where("username =?", username).First(&user)

	if result.Error != nil {
		fmt.Println("no se encontro el usuario ", username)
		return models.User{}, result.Error
	}

	return user, nil
}

func (u *userClient) GetUserById(id int) models.User {
	var user models.User

	result := Db.Where("id =?", id).First(&user)

	if result.Error != nil {
		fmt.Println("no se pudo encontrar el usuario con id: ", id)
		return models.User{}
	}
	return user
}

func (u *userClient) Insert(user models.User) {

	result := Db.Create(&user)

	if result.Error != nil {
		fmt.Printf("Error al cargar el user con nro %v", user.Id)
		return
	}

	fmt.Println("Usuario creado: ", user)
}
