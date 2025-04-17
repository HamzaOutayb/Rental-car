package service

import (
	"errors"
	"strings"

	"social-network/internal/models"
)

func (S *Service) AddHomeinfo(home *models.Home) error {
	var errMsgs []string

	if home.ShopID == 0 {
		errMsgs = append(errMsgs, "shopID is required")
	}

	if home.Name == "" {
		errMsgs = append(errMsgs, "name is required")
	}
	if home.Addr == "" {
		errMsgs = append(errMsgs, "address is required")
	}
	if home.IG == "" {
		errMsgs = append(errMsgs, "Instagram handle is required")
	}
	if home.Facebook == "" {
		errMsgs = append(errMsgs, "Facebook is required")
	}
	if home.Watssapp == "" {
		errMsgs = append(errMsgs, "Whatsapp is required")
	}
	if home.Email == "" {
		errMsgs = append(errMsgs, "email is required")
	}
	if home.Phone_num == "" {
		errMsgs = append(errMsgs, "phone number is required")
	}
	if home.Days_of_week == "" {
		errMsgs = append(errMsgs, "days of week are required")
	}
	if home.Open_time == "" {
		errMsgs = append(errMsgs, "open time is required")
	}
	if home.Close_time == "" {
		errMsgs = append(errMsgs, "close time is required")
	}

	if len(errMsgs) > 0 {
		return errors.New("validation failed: \n" + strings.Join(errMsgs, ",\n"))
	}

	err := S.Database.AddShopInfo(home)
	if err != nil {
		return err
	}

	// continue to save the data
	return nil
}
