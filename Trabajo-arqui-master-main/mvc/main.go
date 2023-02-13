package main

import (
	"mvc/app"
	"mvc/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
