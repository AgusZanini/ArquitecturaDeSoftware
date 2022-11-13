package app

import (
	"time"

	clients "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/clients/rabbitmq"
	controllers "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/controllers"
	service "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/services"
	repositories "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv2/services/repositories"
	//"github.com/rabbitmq/amqp091-go"
)

type Dependencies struct {
	ItemController *controllers.Controller
}

func BuildDependencies() *Dependencies {
	// Repositories
	ccache := repositories.NewCCache(1000, 100, 30*time.Second)
	mongo := repositories.NewMongoDB("localhost", 27017, "items")
	queue := clients.NewRabbitmq("localhost", 5672)

	// Services
	service := service.NewServiceImpl(ccache, mongo, queue)

	// Controllers
	controller := controllers.NewController(service)

	return &Dependencies{
		ItemController: controller,
	}
}
