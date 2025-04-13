package service

import (
	"fmt"
	"social-network/internal/models"
	"strconv"
	"strings"
)

func (s *Service) Addcar(car *models.CarToInsert, imagePaths []string) (int, error) {
	// Insert into cars table
	id, err := s.Database.Addcartodb(car)
	if err != nil {
		return 0, fmt.Errorf("failed to add car to database: %w", err)
	}

	// Set the returned ID to the car object
	car.ID = id

	// Add conditions if they exist
	if len(car.Conditions) > 0 {
		if err := s.Database.AddConditions(car.ID, car.Conditions); err != nil {
			return 0, fmt.Errorf("failed to add conditions: %w", err)
		}
	}

	// Insert images into the database
	for i, imgPath := range imagePaths {
		isPrimary := 0
		if i == 0 {
			isPrimary = 1
		}
		err := s.Database.AddCarImage(car.ID, imgPath, isPrimary)
		if err != nil {
			return 0, fmt.Errorf("failed to add images: %w", err)
		}
	}

	return car.ID, nil
}

func (s *Service) DeleteCar(carID string) error {
	carId, err := strconv.Atoi(carID)
	if err != nil {
		return err
	}

	// delete car from car table
	err = s.Database.DeleteCar(carId)
	if err != nil {
		return err
	}

	// delete images
	err = s.Database.DeleteCarImages(carId)
	if err != nil {
		return err
	}

	// delete car conditions
	err = s.Database.DeleteCarConditions(carId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) EditCar(car *models.CarToEdite) (err error) {
	carIDStr := car.ID
	carID, err := strconv.Atoi(carIDStr)
	if err != nil {
		return
	}

	// insert
	name := car.Name
	description := car.description

	// Handle updated fields
	price, _ := strconv.ParseFloat(car.Price, 64)
	brandID, _ := strconv.Atoi(car.BrandID)
	typeID, _ := strconv.Atoi(car.TypeID)
	contactID, _ := strconv.Atoi(car.ContactID)
	localID, _ := strconv.Atoi(car.LocalID)

	// Extract conditions (comma-separated values)
	conditions := strings.Split(car.Conditions, ",")

	// Optional: handle new images (replace or add)
	if len(car.ImagesToDelete) > 0 && car.ImagesToDelete[0] != "" {
		for _, imagePath := range car.ImagesToDelete {
			err = s.Database.DeleteCarImagesByPath(carID, imagePath)
			if err != nil {
				return
			}
		}
	}

	var isPrimary int

	// add new images
	if len(car.NewImagePaths) > 0 && car.NewImagePaths[0] != "" {
		for i, path := range car.NewImagePaths {
			isPrimary = 0 // default
			if car.Primary == "true" && i == 0 {
				isPrimary = 1 // mark the first one as primary
			}
			err = s.Database.AddCarImage(carID, path, isPrimary)
			if err != nil {
				return
			}
		}
	}
	
	return nil
}
