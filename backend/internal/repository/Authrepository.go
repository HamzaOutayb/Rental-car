package repository

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (data *Repository) CheckMailAndPaswdvalidity(email string, Password string) (int, error) {
	var dbpswd string
	var usrId int
	err := data.Db.QueryRow("SELECT id, password FROM users WHERE email=?", email).Scan(&usrId, &dbpswd)
	if err != nil {
		return 0, errors.New("invalide coredentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbpswd), []byte(Password))
	if err != nil {
		return 0, errors.New("incorrect password")
	}
	return usrId, nil
}

func (database *Repository) AddUuid(Uuid string, userId int) error {
	_, err := database.Db.Exec("INSERT INTO sessions (uuid, user_id, session_exp) VALUES (?,?,?)", Uuid, userId, time.Now().AddDate(1, 0, 0))
	if err != nil {
		return err
	}
	return nil
}
