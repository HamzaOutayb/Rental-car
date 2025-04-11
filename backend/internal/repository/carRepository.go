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
	if err != nil {
		return err
	}

	return nil
}

// I get car condition from db as a slice of byte/json then change it to a slice of string
func (data *Repository) GetConditions(carId int) ([]string, error) {
	var conditionsText string
	err := data.Db.QueryRow("SELECT conditions FROM cars WHERE id = ?", carId).Scan(&conditionsText)
	if err != nil {
		return []string{},err
	}
	var conditions []string
	json.Unmarshal([]byte(conditionsText), &conditions)

	return conditions, nil
}