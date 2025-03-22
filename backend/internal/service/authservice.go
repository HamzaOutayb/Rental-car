package service

import (
	"social-network/internal/models"

	"github.com/gofrs/uuid"
)

func (S *Service) LoginUser(User *models.User) (string, error) {
	// chefk password and email validity
	id, err := S.Database.CheckMailAndPaswdvalidity(User.Email, User.Password)
	if err != nil {
		return "", err
	}

	User.ID = id
	// generate new uuid
	Uuid := GenerateUuid()

	// Update uuid
	S.Database.AddUuid(Uuid, User.ID)
	return Uuid, nil
}

func GenerateUuid() string {
	return uuid.Must(uuid.NewV4()).String()
}
