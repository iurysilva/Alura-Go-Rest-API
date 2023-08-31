package routes

import (
	"log"
	"net/http"

	"github.com/iurysilva/Alura-Go-Rest-API/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
