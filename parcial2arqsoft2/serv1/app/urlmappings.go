package app

import (
	searchcontroller "github.com/AgusZanini/ArquitecturaDeSoftware/parcial2arqsoft2/serv1/controllers"
)

// MapUrls maps the urls
func MapUrls() {

	// Users Mapping
	router.GET("/search=:id", searchcontroller.Search)

}
