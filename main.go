package main

import (
	"fmt"

	"github.com/iurysilva/Alura-Go-Rest-API/database"
	"github.com/iurysilva/Alura-Go-Rest-API/routes"
)

func main() {
	database.ConnectWithDatabase()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
