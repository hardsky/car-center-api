package db

import (
	"github.com/go-pg/pg"
	"github.com/hardsky/car-center-api/models"
	log "github.com/sirupsen/logrus"
)

// DB struct contains methods on top of database.
type DB struct {
	con *pg.DB
}

// Opts contains options for database connection.
type Opts struct {
	Addr     string
	User     string
	Password string
	Database string
}

// Connect eshablishes connection to database,
// returns DB struct with methods on top of database.
func Connect(opt *Opts) *DB {
	return &DB{pg.Connect(&pg.Options{
		Addr:     opt.Addr,
		User:     opt.User,
		Password: opt.Password,
		Database: opt.Database,
	})}
}

// IsDuplicate determines when two users are duplicates.
func (p *DB) IsDuplicate(userID1, userID2 int64) bool {
	var count int
	_, err := p.con.QueryOne(pg.Scan(&count), `
SELECT COUNT (1) FROM
(SELECT ip_addr FROM conn_log WHERE user_id = ?) c1
INNER JOIN
(SELECT ip_addr FROM conn_log WHERE user_id = ?) c2
ON c1.ip_addr = c2.ip_addr
`, userID1, userID2)

	if err != nil {
		log.WithFields(log.Fields{
			"userID1": userID1,
			"userID2": userID2,
		}).WithError(err).Error("check duplicates in db")
	}

	return count >= 2
}

func (p *DB) CreateCar(car *models.Car) (*models.Car, error) {
	return car, p.con.Insert(car)
}

func (p *DB) UpdateCar(car *models.Car) (*models.Car, error) {
	return car, p.con.Update(car)
}

func (p *DB) DeleteCar(id uint64) error {
	return p.con.Delete(&models.Car{SerialNumber: id})
}

func (p *DB) GetCar(id uint64) (*models.Car, error) {
	car := &models.Car{}
	return car, p.con.Model(car).Where("serial_number = ?", id).Select()
}

type PageOpt struct {
	Offset int
	Limit  int
}

func (p *DB) GetCars(opt *PageOpt) ([]*models.Car, error) {
	cars := []*models.Car{}
	return cars, p.con.Model(&cars).Order("serial_number ASC").Limit(opt.Limit).Offset(opt.Offset).Select()
}
