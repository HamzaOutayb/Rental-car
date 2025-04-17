package handler

import (
	"encoding/json"
	"net/http"
	"social-network/internal/models"
	utils "social-network/pkg"
)

/*
multiple forms a footer one a name one aschedule one ext...
*/
func (H *Handler) AddHomeinformations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var Homeinfo models.Home
	err := json.NewDecoder(r.Body).Decode(&Homeinfo)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "encoding the homeinfo")
		return
	}

	// Insert the car into the database
	err = H.Service.AddHomeinfo(&Homeinfo)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, err)
		return
	}

	// Success response
	utils.WriteJson(w, http.StatusCreated, "Home info added successfully")
}

func (H *Handler) GetHomeinformations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	shopID := r.URL.Query().Get("shopID")
	Home, err := H.Service.Database.GetHomeformations(shopID)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, "internal server error")
		return
	}

	utils.WriteJson(w, http.StatusOK, Home)
}
