package handler

import "net/http"

func (H *Handler)  AddCar(w http.ResponseWriter, r *http.Request) {
	r.body 
	H.service.addcar()
}
func (H *Handler)  EditCar(w http.ResponseWriter, r *http.Request) {}
func (H *Handler)  DeleteCar(w http.ResponseWriter, r *http.Request) {}