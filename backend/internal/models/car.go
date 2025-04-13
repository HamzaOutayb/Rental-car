package models

type Car struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       string   `json:"price"`
	Images      []string `json:"images"`
	Avaibility  string   `json:"avaibility"`
	Return_date string   `json:"return_date"`
	Brand       string   `json:"brand"`
	Type        string   `json:"type"`
	Telegram    string   `json:"telegram"`
	Watssapp    string   `json:"watssapp"`
	Main        int      `json:"main"`
	Main_type   int      `json:"main_type"`
}

type CarToInsert struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	BrandID     int      `json:"brand_id"`
	TypeID      int      `json:"type_id"`
	ContactID   int      `json:"contact_id"`
	LocalID     int      `json:"local_id"`
	Conditions  []string `json:"conditions"`
}

type CarToEdite struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	BrandID     string `json:"brand_id"`
	TypeID      string `json:"type_id"`
	ContactID   string `json:"contact_id"`
	LocalID     string `json:"local_id"`
	Conditions  string `json:"conditions"`
}
