package api

import (
	"net/http"

	"github.com/hardsky/car-center-api/db"

	"github.com/gorilla/mux"
)

// Opts contains db.
type Opts struct {
	Db *db.DB
}

// NewAPI constructs new api
func NewAPI(opt *Opts) *API {
	router := mux.NewRouter()
	res := &API{
		d: opt.Db,
		h: router,
	}

	//service routes
	router.HandleFunc("/cars", res.cars).Methods("GET")
	router.HandleFunc("/cars/{carId:[0-9]+}", res.carByID).Methods("GET")
	router.HandleFunc("/cars/add", res.carNew).Methods("POST")
	router.HandleFunc("/cars/{carId:[0-9]+}", res.carEdit).Methods("POST")
	router.HandleFunc("/cars/{carId:[0-9]+}/delete", res.carDelete).Methods("POST")

	return res
}

//API contains methods that implements routes.
type API struct {
	d *db.DB
	h http.Handler
}

// Handler returns http.Handler with api routes.
func (p *API) Handler() http.Handler {
	return p.h
}
