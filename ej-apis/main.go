package main

import (
	"fmt"
	"os"

	"github.com/AgusZanini/ArquitecturaDeSoftware/ej-apis/films"
)

func main() {
	films, err := films.Getfilms("SGA")
	if err != nil {
		fmt.Println("Error getting films: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Pelicula obtenida de api de Studio Ghibli:")
	for _, film := range films {
		fmt.Println(film)
	}
}
