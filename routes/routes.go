package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/iurysilva/Alura-Go-Rest-API/controllers"
	"github.com/iurysilva/Alura-Go-Rest-API/middleware"
)

func HandleRequest() {
	gorillaMux := mux.NewRouter()
	handler := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(gorillaMux)
	gorillaMux.Use(middleware.ContentTypeMiddleware)
	gorillaMux.HandleFunc("/", controllers.Home)
	gorillaMux.HandleFunc("/personalities", controllers.GetAllPersonalities).Methods("Get")
	gorillaMux.HandleFunc("/personalities/{id}", controllers.GetOnePersonality).Methods("Get")
	gorillaMux.HandleFunc("/personalities", controllers.CreateNewPersonality).Methods("Post")
	gorillaMux.HandleFunc("/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	gorillaMux.HandleFunc("/personalities/{id}", controllers.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
