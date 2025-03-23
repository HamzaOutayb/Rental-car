package service

import "social-network/internal/models"

func (s *Service) Addcar(car *models.CarToInsert) (int64, error){
	//inset in cars tables

	// add conditions
	AddConditions(0, car.Conditions)
	return 0, nil
}


func (s *Service) AddCarImage(carId int64, path string, isPrimary int) (error){
	
	return nil
}

func AddConditions(carId int64, consitions []string) (error){
	
	return  nil
}