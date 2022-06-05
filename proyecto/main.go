package main

import(
	db "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/db"
	app "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/app"
)

func main() {
	app.StartRoutes()
	db.StartDbEngine()
}