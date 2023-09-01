package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
