package main

import (
	"fmt"
	"os"

	"github.com/AgusZanini/ArquitecturaDeSoftware/ej-ml/api"
)

func main() {
	categories, err := api.Getcategories("MLA")
	if err != nil {
		fmt.Println("Error al obtener las categorias: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Categoria obtenida: ")
	for _, category := range categories {
		fmt.Println(category.String())
	}
}
