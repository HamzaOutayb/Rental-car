package service

import (
	"fmt"
	"social-network/internal/models"
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
		err := s.Database.AddCarImage(car.ID, imgPath, isPrimary); if err != nil {
			return 0, fmt.Errorf("failed to add images: %w", err)
		}
	}

	return car.ID, nil
}
