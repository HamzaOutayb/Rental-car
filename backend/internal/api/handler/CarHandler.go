package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"social-network/internal/models"
	utils "social-network/pkg"
	"strconv"
	"strings"
)

const uploadDir = "repository/uploads"

func (H *Handler) AddCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	r.ParseMultipartForm(10 << 20) // 10MB max

	// Handle image uploads
	files := r.MultipartForm.File["images"]
	var imagePaths []string

	for i, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			utils.WriteJson(w, http.StatusBadRequest, "Failed to open image file")
			return
		}
		defer file.Close()
		Uuid := utils.GenerateUuid()
		// Save file to server storage
		imagePath := fmt.Sprintf("uploads/%s%s_%d%s",uploadDir, Uuid, i, filepath.Ext(fileHeader.Filename))

		dst, err := os.Create(imagePath)
		if err != nil {
			utils.WriteJson(w, http.StatusInternalServerError, "Failed to save image")
			return
		}
		defer dst.Close()
		io.Copy(dst, file)

		imagePaths = append(imagePaths, imagePath)
	}

	// Extract car details from form data
	name := r.FormValue("name")
	description := r.FormValue("description")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	brandID, _ := strconv.Atoi(r.FormValue("brand_id"))
	typeID, _ := strconv.Atoi(r.FormValue("type_id"))
	contactID, _ := strconv.Atoi(r.FormValue("contact_id"))
	localID, _ := strconv.Atoi(r.FormValue("local_id"))
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
	carID, err := H.Service.Addcar(&car, imagePaths)
	if err != nil {
		http.Error(w, "Failed to add car", http.StatusInternalServerError)
		return
	}

	// Success response
	utils.WriteJson(w, http.StatusCreated, map[string]interface{}{
		"message": "Car added successfully",
		"car_id":  carID,
	})
}

func (H *Handler) EditCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	r.ParseMultipartForm(10 << 20)

	// Handle updated fields
	carID := r.FormValue("car_id")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	brandID := r.FormValue("brand_id")
	typeID := r.FormValue("type_id")
	contactID := r.FormValue("contact_id")
	localID := r.FormValue("local_id")
	conditions := r.FormValue("conditions")

	// check if there is any images to delete
	imagesToDelete := strings.Split(r.FormValue("delete_images"), ",")
	primary := r.FormValue("primary") // primaryImage

	// Optional: handle new images
	files := r.MultipartForm.File["images"]
	var newImagePaths []string

	for i, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			utils.WriteJson(w, http.StatusBadRequest, "Failed to open image file")
			return
		}
		defer file.Close()

		Uuid := utils.GenerateUuid()
		imagePath := fmt.Sprintf("uploads/%s%s_%d%s", uploadDir, Uuid, i, filepath.Ext(fileHeader.Filename))

		dst, err := os.Create(imagePath)
		if err != nil {
			utils.WriteJson(w, http.StatusInternalServerError, "Failed to save image")
			return
		}
		defer dst.Close()
		io.Copy(dst, file)

		newImagePaths = append(newImagePaths, imagePath)
	}

	updatedCar := models.CarToEdite{
		ID:             carID,
		Name:           name,
		Description:    description,
		Price:          price,
		BrandID:        brandID,
		TypeID:         typeID,
		ContactID:      contactID,
		LocalID:        localID,
		Conditions:     conditions,
		NewImagePaths:  newImagePaths,
		ImagesToDelete: imagesToDelete,
		Primary:        primary,
	}

	err := H.Service.EditCar(&updatedCar)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "failed to update car")
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"message": "Car updated successfully",
		"car_id":  carID,
	})
}

func (H *Handler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	carId := r.FormValue("cardId")
	err := H.Service.DeleteCar(carId)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "err while deleting a car")
		return
	}

	utils.WriteJson(w, http.StatusOK, "cra deleted successfuly")
}

func (H *Handler) GetCarsbyBrand(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	brandName, page := r.URL.Query().Get("brand"), r.URL.Query().Get("page")
	Cars, err := H.Service.GetCarByBrandID(brandName, page)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "method not allowed")
		return
	}

	utils.WriteJson(w, http.StatusOK, Cars)
}

/*getting car by Type ex:*/
/*Type: luxury, SUv , Sedan, supercar, van, economie, electric, busness, convertebel*/
func (H *Handler) GetCarsbyType(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	TypeID, page := r.URL.Query().Get("type"), r.URL.Query().Get("page")
	Cars, err := H.Service.GetCarByTypeID(TypeID, page)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, "method not allowed")
		return
	}

	utils.WriteJson(w, http.StatusOK, Cars)
}

func (H *Handler) GetTrending(w http.ResponseWriter, r *http.Request)   {}
func (H *Handler) GetTopRentals(w http.ResponseWriter, r *http.Request) {}

func (H *Handler) Getcarbyid(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		utils.WriteJson(w, http.StatusBadRequest, "Missing car ID")
		return
	}

	Car, err := H.Service.Database.GetCarByID(id)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, Car)
}
