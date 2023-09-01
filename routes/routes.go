package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iurysilva/Alura-Go-Rest-API/controllers"
)

func HandleRequest() {
	gorillaMux := mux.NewRouter()
	gorillaMux.HandleFunc("/", controllers.Home)
	gorillaMux.HandleFunc("/getAllPersonalities", controllers.GetAllPersonalities)
	gorillaMux.HandleFunc("/getOnePersonality/{id}", controllers.GetOnePersonality)
	log.Fatal(http.ListenAndServe(":8000", gorillaMux))
}
