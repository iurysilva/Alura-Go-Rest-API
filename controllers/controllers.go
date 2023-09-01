package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iurysilva/Alura-Go-Rest-API/database"
	"github.com/iurysilva/Alura-Go-Rest-API/models"
)

func Home(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "Home Page")
}

func GetAllPersonalities(responseWriter http.ResponseWriter, request *http.Request) {
	var personalities []models.Personality
	database.DB.Find(&personalities)
	json.NewEncoder(responseWriter).Encode(personalities)
}

func GetOnePersonality(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	var personality models.Personality
	database.DB.First(&personality, "id = ?", id)
	json.NewEncoder(responseWriter).Encode(personality)
}

func CreateNewPersonality(responseWriter http.ResponseWriter, request *http.Request) {
	var personality models.Personality
	json.NewDecoder(request.Body).Decode(&personality)
	database.DB.Create(&personality)
	json.NewEncoder(responseWriter).Encode(personality)
}
