package api

import (
	"database/sql"
	"net/http"
	"social-network/internal/api/handler"
)

func Routes(db *sql.DB) *http.ServeMux {
	handler := handler.NewHandler(db)
	mux := http.NewServeMux()

	// login
	mux.HandleFunc("/api/login", handler.Login)


	// dashboard //
	mux.HandleFunc("/api/addcar", handler.AddCar)
	mux.HandleFunc("/api/editCar", handler.EditCar)
	mux.HandleFunc("/api/deleteCar", handler.DeleteCar)

	// local //
/*	mux.HandleFunc("/api/createlocal", handler.Createlocal)
	mux.HandleFunc("/api/updatelocal", handler.Updatelocal)
	mux.HandleFunc("/api/deletelocal", handler.Deletelocal)
*/

	// Home page //
	mux.HandleFunc("/api/getHome/", handler.GetHomeinformations) // home footerinfo name schedule ext...
	mux.HandleFunc("/api/addHome/", handler.AddHomeinformations) // home footerinfo name schedule ext...
	// Getting cars //
	mux.HandleFunc("/api/getcars/brand/", handler.GetCarsbyBrand)// /api/getcars/barnd/BMW
				/*Getting cars by type ID*/
	//Type: luxury, SUv , Sedan, supercar, van, economie, electric, busness, convertebel //
	mux.HandleFunc("/api/getcars/type/", handler.GetCarsbyType)
/*	mux.HandleFunc("/api/trending/", handler.GetTrending)
	mux.HandleFunc("/api/top-rentals/", handler.GetTopRentals)*/	
	mux.HandleFunc("/api/getcar/{id}", handler.Getcarbyid)


	// Like/register functionality
	/*mux.HandleFunc("/api/LikeCar", handler.Getcarbyid)
	mux.HandleFunc("/api/LikedCars", handler.GetLikedCars)*/

	return mux
}