package main

import (
	"proyectoarqsoft/app"
	"proyectoarqsoft/db"
)

func main() {
	db.Connect()
	app.StartRoute()
}
