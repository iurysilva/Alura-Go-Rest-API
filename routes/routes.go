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
	gorillaMux.HandleFunc("/personalities", controllers.GetAllPersonalities).Methods("Get")
	gorillaMux.HandleFunc("/personalities/{id}", controllers.GetOnePersonality).Methods("Get")
	gorillaMux.HandleFunc("/personalities", controllers.CreateNewPersonality).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", gorillaMux))
}
