package repository

import (
	"encoding/json"
	"errors"
	"social-network/internal/models"
)

func (data *Repository) Addcartodb(car *models.CarToInsert) (int, error) {
	row, err := data.Db.Exec("INSERT INTO cars (name, description, price, brand_id, type_id, contact_id) VALUES (?, ?, ?, ?, ?, ?)",
		&car.Name,
		&car.Description,
		&car.Price,
		&car.BrandID,
		&car.TypeID,
		&car.ContactID,
	)

	if err != nil {
		return 0, errors.New("error while inserting the car data")
	}

	id, err := row.LastInsertId()
	if err != nil {
		return 0, errors.New("error while getting the last inserted id")
	}

	return int(id), nil
}

func (data *Repository) AddConditions(carId int, conditions []string) error {
	_, err := data.Db.Exec("INSERT INTO conditions (condition, car_id)", carId, conditions)
	return err
}

// I get car condition from db as a slice of byte/json then change it to a slice of string
func (data *Repository) GetConditions(carId int) ([]string, error) {
	var conditionsText string
	err := data.Db.QueryRow("SELECT conditions FROM cars WHERE id = ?", carId).Scan(&conditionsText)
	if err != nil {
		return []string{}, err
	}
	var conditions []string
	json.Unmarshal([]byte(conditionsText), &conditions)

	return conditions, nil
}

func (data *Repository) AddCarImage(carId int, path string, isPrimary int) error {
	_, err := data.Db.Exec("INSERT INTO car_images (car_id, image_path, is_primary)", carId, path, isPrimary)
	return err
}

func (data *Repository) DeleteCar(carId int) error {
	_, err := data.Db.Exec("DELETE FROM cars WHERE id = ?", carId)
	return err
}

func (data *Repository) DeleteCarImages(carId int) error {
	_, err := data.Db.Exec("DELETE FROM car_images WHERE car_id = ?", carId)
	return err
}

func (data *Repository) DeleteCarImagesByPath(carId int, image string) error {
	_, err := data.Db.Exec("DELETE FROM car_images WHERE car_id = ? AND image_path = ?", carId, image)
	return err
}

func (data *Repository) DeleteCarConditions(carId int) error {
	_, err := data.Db.Exec("DELETE FROM car_conditions WHERE car_id = ?", carId)
	return err
}

func (data *Repository) CarToEdite(car *models.CarToInsert) error {
	err := data.Db.Exec("UPDATE")
}