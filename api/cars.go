package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hardsky/car-center-api/db"
	"github.com/hardsky/car-center-api/models"
	log "github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type CarResponse struct {
	Car *models.Car `json:"car"`
}

// get list
func (p *API) cars(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	limit, err := strconv.ParseInt(vars["pageSize"], 10, 64)
	if err != nil || limit <= 0 {
		log.WithError(err).Error("wrong limit")
		writeResponse(w, &ErrorResponse{"wrong limit"}, http.StatusBadRequest)
		return
	}

	offset, err := strconv.ParseInt(vars["page"], 10, 64)
	offset -= 1
	if err != nil || offset < 0 {
		log.WithError(err).Error("wrong offset")
		writeResponse(w, &ErrorResponse{"wrong page"}, http.StatusBadRequest)
		return
	}

	cars, err := p.d.GetCars(&db.PageOpt{Offset: int(offset), Limit: int(limit)})
	if err != nil {
		log.WithError(err).Error("db GetCars")
		writeResponse(w, &ErrorResponse{"internal server error"}, http.StatusBadRequest)
		return
	}

	writeResponse(w, cars, http.StatusOK)
}

// get record
func (p *API) carByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctxLog := log.WithFields(log.Fields{
		"carID": vars["carId"],
	})

	carID, err := strconv.ParseInt(vars["carID"], 10, 64)
	if err != nil || carID <= 0 {
		ctxLog.WithError(err).Error("wrong car Id")
		writeResponse(w, &ErrorResponse{"wrong car Id"}, http.StatusBadRequest)
		return
	}

	car, err := p.d.GetCar(uint64(carID))
	if err != nil {
		ctxLog.WithError(err).Error("db GetCar")
		writeResponse(w, &ErrorResponse{"car is not found"}, http.StatusBadRequest)
		return
	}

	writeResponse(w, car, http.StatusOK)
}

// add new record
func (p *API) carNew(w http.ResponseWriter, r *http.Request) {
	car := &models.Car{}
	err := json.NewDecoder(r.Body).Decode(car)
	if err != nil {
		log.WithError(err).Error("parse new car request")
		writeResponse(w, &ErrorResponse{"error in json"}, http.StatusBadRequest)
		return
	}

	car, err = p.d.CreateCar(car)
	if err != nil {
		log.WithError(err).Error("db new car")
		writeResponse(w, &ErrorResponse{"internal error"}, http.StatusInternalServerError)
		return
	}

	writeResponse(w, car, http.StatusOK)
}

// edit record
func (p *API) carEdit(w http.ResponseWriter, r *http.Request) {
	car := &models.Car{}
	err := json.NewDecoder(r.Body).Decode(car)
	if err != nil {
		log.WithError(err).Error("parse update car request")
		writeResponse(w, &ErrorResponse{"error in json"}, http.StatusBadRequest)
		return
	}

	car, err = p.d.UpdateCar(car)
	if err != nil {
		log.WithError(err).Error("db update car")
		writeResponse(w, &ErrorResponse{"internal error"}, http.StatusInternalServerError)
		return
	}

	writeResponse(w, car, http.StatusOK)
}

// delete record
func (p *API) carDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctxLog := log.WithFields(log.Fields{
		"carID": vars["carId"],
	})

	carID, err := strconv.ParseInt(vars["carID"], 10, 64)
	if err != nil || carID <= 0 {
		ctxLog.WithError(err).Error("wrong car Id")
		writeResponse(w, &ErrorResponse{"wrong car Id"}, http.StatusBadRequest)
		return
	}

	err = p.d.DeleteCar(uint64(carID))
	if err != nil {
		ctxLog.WithError(err).Error("db delete car")
		writeResponse(w, &ErrorResponse{"car is not found"}, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func writeResponse(w http.ResponseWriter, body interface{}, status int) {
	b, _ := json.Marshal(body)
	w.WriteHeader(status)
	w.Write(b)
	w.Header().Set("Content-Type", "application/json")
}
