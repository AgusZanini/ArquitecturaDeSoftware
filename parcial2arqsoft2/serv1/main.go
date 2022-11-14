package main

import (
	app "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/app"
	clients "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/clients"
)

func main() {
	go clients.Consumer()
	app.StartRoute()

}
