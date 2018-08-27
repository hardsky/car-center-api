package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating tables...")
		_, err := db.Exec(`
CREATE TABLE accelerations (
        id BIGSERIAL PRIMARY KEY,
        mph BIGINT,
        seconds double precision
        );

CREATE TABLE performance_figures (
        id BIGSERIAL PRIMARY KEY,
        octane_rating INT,
        acceleration_id BIGINT REFERENCES accelerations ON DELETE SET NULL
        );

CREATE TABLE engines (
        id BIGSERIAL PRIMARY KEY,
        capacity INT,
        num_cylinders INT,
        max_rpm INT,
        manufacturer_code TEXT
        );

CREATE TABLE fuel_figures (
        id BIGSERIAL PRIMARY KEY,
        speed INT,
        mpg DOUBLE PRECISION,
        usage_description TEXT
        );

CREATE TABLE cars (
        serial_number BIGINT PRIMARY KEY,
        owner_name TEXT,
        model_year BIGINT,
        code TEXT,
        vehicle_code TEXT,
        manufacturer TEXT,
        model TEXT,
        activation_code TEXT,
        engine_id BIGINT REFERENCES engines ON DELETE SET NULL,
        fuel_figure_id BIGINT REFERENCES fuel_figures ON DELETE SET NULL,
        performance_figure_id BIGINT REFERENCES performance_figures ON DELETE SET NULL
        );
`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping tables...")
		_, err := db.Exec(`
DROP TABLE cars;
DROP TABLE performance_figures;
DROP TABLE engines;
DROP TABLE fuel_figures;
DROP TABLE accelerations;
`)
		return err
	})
}
