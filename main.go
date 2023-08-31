package main

import (
	"fmt"

	"github.com/iurysilva/Alura-Go-Rest-API/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
