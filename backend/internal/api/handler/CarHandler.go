package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"social-network/internal/models"
)

func (H *Handler) AddCar(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10MB max

	// Extract car details from form data
	name := r.FormValue("name")
	description := r.FormValue("description")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	brandID, _ := strconv.Atoi(r.FormValue("brand_id"))
	typeID, _ := strconv.Atoi(r.FormValue("type_id"))
	contactID, _ := strconv.Atoi(r.FormValue("contact_id"))
	localID, _ := strconv.Atoi(r.FormValue("local_id"))

	// Extract conditions (comma-separated values)
	conditions := strings.Split(r.FormValue("conditions"), ",")

	// Create a car object
	car := models.CarToInsert{
		Name:        name,
		Description: description,
		Price:       price,
		BrandID:     brandID,
		TypeID:      typeID,
		ContactID:   contactID,
		LocalID:     localID,
		Conditions:  conditions,
	}

	// Insert the car into the database
	carID, err := H.Service.Addcar(&car)
	if err != nil {
		http.Error(w, "Failed to add car", http.StatusInternalServerError)
		return
	}

	// Handle image uploads
	files := r.MultipartForm.File["images"]
	var imagePaths []string

	for i, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Failed to open image file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Save file to server storage
		imagePath := fmt.Sprintf("uploads/%d_%d%s", carID, i, filepath.Ext(fileHeader.Filename))
		dst, err := os.Create(imagePath)
		if err != nil {
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		io.Copy(dst, file)

		imagePaths = append(imagePaths, imagePath)
	}

	// Insert images into the database
	for i, imgPath := range imagePaths {
		isPrimary := 0
		if i == 0 {
			isPrimary = 0
		}
		H.Service.AddCarImage(carID, imgPath, isPrimary)
	}

	// Success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Car added successfully",
		"car_id":  carID,
	})
}
func (H *Handler) EditCar(w http.ResponseWriter, r *http.Request)   {}
func (H *Handler) DeleteCar(w http.ResponseWriter, r *http.Request) {}
