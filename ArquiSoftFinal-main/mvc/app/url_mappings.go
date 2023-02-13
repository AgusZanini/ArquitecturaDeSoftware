package app

import (
	categoryController "mvc/controllers/category"
	orderController "mvc/controllers/order"
	orderDetailController "mvc/controllers/order_detail"
	productController "mvc/controllers/product"
	userController "mvc/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/product/:product_id", productController.GetProductById)
	router.GET("/products", productController.GetProducts)
	router.GET("/products/:category_id", productController.GetProductsByCategoryId)
	router.GET("/products/search=:searchQuery", productController.GetProductsBySearch)

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert)

	// Orders Mapping
	router.GET("/order/:id", orderController.GetOrderById)
	router.GET("/orders/:id", orderController.GetOrdersByUserId)
	router.POST("/order", orderController.OrderInsert)

	// OrderDetails Mapping
	router.GET("/orderDetail/:id", orderDetailController.GetOrderDetailById)
	router.POST("/orderDetail", orderDetailController.OrderDetailInsert)

	// Category Mapping
	router.GET("/category/:id", categoryController.GetCategoryById)
	router.GET("/categories", categoryController.GetCategories)

	// Login Mapping
	router.POST("/login", userController.Login)

	log.Info("Finishing mappings configurations")
}
