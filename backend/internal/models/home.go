package models

type Home struct {
	ShopID    int               `json:"shopid"`
	Brands    map[string]string // kay:brandname value:imagepath
	Name      string            `json:"name"`
	Addr      string            `json:"addr"`
	IG        string            `json:"ig"`
	Facebook  string            `json:"facebook"`
	Watssapp  string            `json:"watssapp"`
	Email     string            `json:"email"`
	Phone_num string            `json:"phone_num"`
	// schedule
	Days_of_week string `json:"daysofweek"`
	Open_time    string `json:"openTime"`
	Close_time   string `json:"closeTime"`
}