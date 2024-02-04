package porscheapi

import (
	"errors"
)

type Car struct {
	ID        string   `json:"id"`
	Brand     string `json:"brand"`
	Model     string `json:"model"`
	Year      string   `json:"year"`
}

var cars = []Car{
	{
		ID:      "1",
		Brand:   "Porsche",
		Model:   "911 Turbo S",
		Year:    "2022",
	},
	{
		ID:      "2",
		Brand:   "Porsche",
		Model:   "Cayman GT4",
		Year:    "2023",
	},
	{
		ID:      "3",
		Brand:   "Porsche",
		Model:   "Panamera Turbo S E-Hybrid",
		Year:    "2022",
	},
	{
		ID:      "4",
		Brand:   "Porsche",
		Model:   "Cayenne",
		Year:    "2022",
	},
	{
		ID:      "5",
		Brand:   "Porsche",
		Model:   "Macan GTS",
		Year:    "2023",
	},
	{
		ID:      "6",
		Brand:   "Porsche",
		Model:   "Boxstar",
		Year:    "2022",
	},
}

func GetCars() []Car {
	return cars
}

func GetCar(id string) (*Car, error) {
	for _, r := range cars {
		if r.ID == id {
			return &r, nil
		}
	}
	return nil, errors.New("porsche not found")
}