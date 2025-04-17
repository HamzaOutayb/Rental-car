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

func (data *Repository) CarToEdite(car *models.CarToEdite) error {
	_, err := data.Db.Exec(`
	UPDATE cars
	SET
		name = CASE WHEN ? != '' THEN ? ELSE name END,
		description = CASE WHEN ? != '' THEN ? ELSE description END,
		price = CASE WHEN ? != '' THEN ? ELSE price END,
		brand_id = CASE WHEN ? != '0' THEN ? ELSE brand_id END,
		type_id = CASE WHEN ? != '0' THEN ? ELSE type_id END,
		contact_id = CASE WHEN ? != '0' THEN ? ELSE contact_id END,
		local_id = CASE WHEN ? != '0' THEN ? ELSE local_id END
	WHERE id = ?
	`,
		&car.Name, &car.Name,
		&car.Description, &car.Description,
		&car.Price, &car.Price,
		&car.BrandID, &car.BrandID,
		&car.TypeID, &car.TypeID,
		&car.ContactID, &car.ContactID,
		&car.LocalID, &car.LocalID,
		&car.ID, &car.ID,
	)
	return err
}

func (data *Repository) GetCarsCount(kind, Id string) (int, error) {
	var count int
	switch kind {
	case "type":
		err := data.Db.QueryRow("SELECT FROM type WHERE id = ?", Id).Scan(&count)
		return count, err
	case "brand":
		err := data.Db.QueryRow("DELETE FROM brands WHERE id = ?", Id).Scan(&count)
		return count, err
	}
	return 0, errors.New("GetCarsCount")
}

const LIMIT = 15

func (data *Repository) GetCarbyBrandID(BrandID string, start int) ([]models.Car, error) {
	rows, err := data.Db.Query("SELECT id, name, price, availability_status FROM cars WHERE Brand_id = ? OFFSET ? LIMIT ?", BrandID, start, LIMIT)
	if err != nil {
		return []models.Car{}, err
	}

	var Cars []models.Car
	for rows.Next() {
		var Car models.Car
		err := rows.Scan(
			&Car.ID,
			&Car.Name,
			&Car.Price,
			&Car.Avaibility,
		)
		if err != nil {
			return []models.Car{}, err
		}
		Cars = append(Cars, Car)
	}

	return Cars, nil
}

func (data *Repository) GetCarByTypeID(TypeID string, start int) ([]models.Car, error) {
	rows, err := data.Db.Query("SELECT id, name, price, availability_status FROM cars WHERE type_id = ? OFFSET ? LIMIT ?", TypeID, start, LIMIT)
	if err != nil {
		return []models.Car{}, err
	}

	var Cars []models.Car
	for rows.Next() {
		var Car models.Car
		err := rows.Scan(
			&Car.ID,
			&Car.Name,
			&Car.Price,
			&Car.Avaibility,
		)
		if err != nil {
			return []models.Car{}, err
		}
		Cars = append(Cars, Car)
	}

	return Cars, nil
}

func (data *Repository) GetCarByID(id string) (models.Car, error) {
	var car models.Car
	query := `SELECT id,
	car_id,
	car_name,
	car_description,
	car_price,
	availability_status,
	brand_name,
	type_name,
	contact_name,
	contact_Telegram,
	contact_watssapp,
	FROM cars_detailed_view WHERE id = ?`

	err := data.Db.QueryRow(query, id).Scan(
		&car.ID, &car.Name, &car.Description, &car.Price, &car.Avaibility,
		&car.Brand, &car.Type, &car.Contact_name, &car.Telegram, &car.Watssapp,
	)
	if err != nil {
		return models.Car{}, err
	}

	return car, nil
}
