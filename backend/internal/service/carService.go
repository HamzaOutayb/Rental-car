package service

import "html"

func (s *Service) Addcar(car *models.Car) error {
	car = html.EscapeString(car.content)

	err := s.Database.insertCar(car)
}