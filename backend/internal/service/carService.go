package service

import (
	"errors"
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
	var carInsert models.CarToInsert
	carID, err := strconv.Atoi(car.ID)
	if err == nil {
		carInsert.ID = carID
	}

	if car.Name != "" {
		carInsert.Name = car.Name
	}

	if car.Description != "" {
		carInsert.Description = car.Description
	}

	if price, err := strconv.ParseFloat(car.Price, 64); err == nil {
		carInsert.Price = price
	}

	if brandID, err := strconv.Atoi(car.BrandID); err == nil {
		carInsert.BrandID = brandID
	}

	if typeID, err := strconv.Atoi(car.TypeID); err == nil {
		carInsert.TypeID = typeID
	}

	if contactID, err := strconv.Atoi(car.ContactID); err == nil {
		carInsert.ContactID = contactID
	}

	if localID, err := strconv.Atoi(car.LocalID); err == nil {
		carInsert.LocalID = localID
	}

	if car.Conditions != "" {
		carInsert.Conditions = strings.Split(car.Conditions, ",")
		err = s.Database.AddConditions(carID, carInsert.Conditions)
		if err != nil {
			return err
		}
	}

	err = s.Database.CarToEdite(car)
	if err != nil {
		return err
	}

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

func (s *Service) GetCarByBrandID(brandID, page string) ([]models.Car, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		return nil, errors.New("invalid page number")
	}

	start := (pageNum - 1) * 15

	count, err := s.Database.GetCarsCount("brand", brandID)
	if err != nil {
		return nil, errors.New("error while counting cars by brand")
	}

	if start >= count {
		return nil, errors.New("page out of range")
	}

	cars, err := s.Database.GetCarbyBrandID(brandID, start)
	if err != nil {
		return nil, errors.New("failed to fetch cars by brand")
	}

	return cars, nil
}

func (s *Service) GetCarByTypeID(typeID, page string) ([]models.Car, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		return nil, errors.New("invalid page number")
	}

	start := (pageNum - 1) * 15

	count, err := s.Database.GetCarsCount("type", typeID)
	if err != nil {
		return nil, errors.New("error while counting cars by type")
	}

	if start >= count {
		return nil, errors.New("page out of range")
	}

	cars, err := s.Database.GetCarByTypeID(typeID, start)
	if err != nil {
		return nil, errors.New("failed to fetch cars by type")
	}

	return cars, nil
}
