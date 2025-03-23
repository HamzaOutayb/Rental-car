package models

type Car struct {
	ID int `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    Price string `json:"price"`
    Images []string `json:"images"`
    Avaibility string `json:"avaibility"`
    Return_date string `json:"return_date"`
    Brand string `json:"brand"`
    Type string `json:"type"`
    Telegram string `json:"telegram"`
	Watssapp string `json:"watssapp"`
    Main int `json:"main"`
    Main_type int `json:"main_type"`
}