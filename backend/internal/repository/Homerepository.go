package repository

import "social-network/internal/models"

func (data *Repository) GetHomeformations(shopID string) ([]models.Home, error) {
	/*
		create a view that has
		brand avaible / brands where count != 0
		get all shop info by id
		and last and not least schedule by shop id
	*/
	var Homeinfo []models.Home
	rows, err := data.Db.Query("SELECT * FROM homeview WHERE shopID = ?", shopID)
	if err != nil {
		return []models.Home{}, err
	}
	for rows.Next() {
		var home models.Home
		err := rows.Scan(
			&home.ShopID,
			&home.Brands,
			&home.Name,
			&home.Addr,
			&home.IG,
			&home.Facebook,
			&home.Watssapp,
			&home.Email,
			&home.Phone_num,
			&home.Days_of_week,
			&home.Open_time,
			&home.Close_time,
		)
		if err != nil {
			return []models.Home{}, err
		}
		Homeinfo = append(Homeinfo, home)
	}

	return Homeinfo, nil
}
