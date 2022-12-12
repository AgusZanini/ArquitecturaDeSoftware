package main

import (
	app "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/app"
	db "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/db"
)

func main() {
	app.StartRoutes()
	db.StartDbEngine()
}
