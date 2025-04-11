package service

import (
	"fmt"
	"social-network/internal/models"
)

func (s *Service) Addcar(car *models.CarToInsert) (int, error) {
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

	return car.ID, nil
}
func (s *Service) AddCarImage(carId int, path string, isPrimary int) error {
	
	return nil
}