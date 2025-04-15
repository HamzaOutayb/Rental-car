package handler

import (
	"net/http"
	utils "social-network/pkg"
)

func (H *Handler) GetHomeinformations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	shopID := r.URL.Query().Get("shopID")
	Home,err := H.Service.Database.GetHomeformations(shopID); if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, "internal server error")
		return
	}

	utils.WriteJson(w, http.StatusOK, Home)
}
