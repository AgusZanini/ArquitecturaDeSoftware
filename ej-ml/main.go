package main

import (
	"fmt"
	"os"
	"github.com/AgusZanini/ArquitecturaDeSoftware/ej-ml/api"
)

func main() {
	categories, err := api.getcategories("MLA")
	if err != nil {
		fmt.Println("Error al obtener las categorias: ", err.Error())
		os.Exit(1)
	}
	for _, category != range categories {
		fmt.Println(category.String())
	}
}
