package service

import (
	"social-network/internal/models"
	utils "social-network/pkg"
)

func (S *Service) LoginUser(User *models.User) (string, error) {
	// chefk password and email validity
	id, err := S.Database.CheckMailAndPaswdvalidity(User.Email, User.Password)
	if err != nil {
		return "", err
	}

	User.ID = id
	// generate new uuid
	Uuid := utils.GenerateUuid()

	// Update uuid
	S.Database.AddUuid(Uuid, User.ID)
	return Uuid, nil
}
