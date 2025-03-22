package handler

import (
	"fmt"
	"net/http"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (H *Handler) Login(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var user models.User
	if err := utils.ParseBody(r, &user); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "Bad request")
		return
	}

	Uuid, err := H.Service.LoginUser(&user)
	if err != nil {
		fmt.Println("err", err)
		utils.WriteJson(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SetSessionCookie(w, Uuid)
	utils.WriteJson(w, http.StatusOK, "you'v loged in succesfuly!")
}
