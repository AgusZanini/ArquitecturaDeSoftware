package main

import ("fmt"
		"time")

func main() {
	go process()

	time.Sleep(1 * time.Second)
}

func process() {
	fmt.Println("Procesando 1... ")
	fmt.Println("Procesando 2... ")
}

