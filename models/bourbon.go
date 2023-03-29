package models

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Bourbon struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Size        string  `json:"size"`
	Price       float64 `json:"price"`
	Abv         float64 `json:"abv"`
	Description string  `json:"description"`
}

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./bourbon.sqlite3")

	if err != nil {
		return err
	}

	DB = db
	return nil
}

func GetBourbons(count int) ([]Bourbon, error) {

	rows, err := DB.Query("SELECT * from bourbons LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	bourbons := make([]Bourbon, 0)

	for rows.Next() {
		singleBourbon := Bourbon{}
		err = rows.Scan(&singleBourbon.Id, &singleBourbon.Name, &singleBourbon.Size, &singleBourbon.Price, &singleBourbon.Abv, &singleBourbon.Description)

		if err != nil {
			return nil, err
		}

		bourbons = append(bourbons, singleBourbon)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return bourbons, err
}
