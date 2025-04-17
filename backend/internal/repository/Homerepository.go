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

func (data *Repository) AddShopInfo(home *models.Home) error {
	tx, err := data.Db.Begin()
	if err != nil {
		return err
	}

	// 1. Insert into schedule
	scheduleStmt, err := tx.Prepare(`INSERT INTO schedule (day_of_week, open_time, close_time) VALUES (?, ?, ?)`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer scheduleStmt.Close()

	res, err := scheduleStmt.Exec(home.Days_of_week, home.Open_time, home.Close_time)
	if err != nil {
		tx.Rollback()
		return err
	}
	scheduleID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	// 2. Insert into shop_info
	shopStmt, err := tx.Prepare(`
		INSERT INTO shop_info (name, addr, IG, facebook, watssapp, email, phone_num, schedule_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer shopStmt.Close()

	_, err = shopStmt.Exec(
		home.Name,
		home.Addr,
		home.IG,
		home.Facebook,
		home.Watssapp,
		home.Email,
		home.Phone_num,
		scheduleID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
