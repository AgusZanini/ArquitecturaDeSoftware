package app

import (
	categoryhandler "proyectoarqsoft/controllers/categories"
	orderhandler "proyectoarqsoft/controllers/orders"
	producthandler "proyectoarqsoft/controllers/products"
	userhandler "proyectoarqsoft/controllers/users"
)

func mapUrls() {

	// users mappings

	router.POST("/users", userhandler.Insert)
	router.GET("/users/:id", userhandler.GetUserById)
	router.GET("/login", userhandler.Login)

	// products mappings

	router.POST("/products", producthandler.InsertProduct)
	router.GET("/products/search=:Query", producthandler.SearchByName)
	router.GET("/products/category/:name", producthandler.GetProductsByCategory)

	// categories mappings

	router.POST("/categories", categoryhandler.InsertCategory)

	// orders mappings

	router.POST("/orders", orderhandler.InsertOrder)
	router.GET("/users/orders/:userid", orderhandler.GetOrdersByUser)
}
