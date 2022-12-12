package main

import (
	app "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/app"
	db "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/db"
)

func main() {
	app.StartRoutes()
	db.StartDbEngine()
	db.PopulateDb()
}
