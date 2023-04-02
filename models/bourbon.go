package models

import (
	"database/sql"
	"encoding/json"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Bourbon struct {
	Id          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Size        string  `json:"size,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Abv         float64 `json:"abv,omitempty"`
	Description string  `json:"description,omitempty"`
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

	rows, err := DB.Query("SELECT id, name, size, price, abv, description from bourbons LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	bourbons := make([]Bourbon, 0)

	for rows.Next() {
		singleBourbon := Bourbon{}
		err = rows.Scan(
			&singleBourbon.Id,
			&singleBourbon.Name,
			&singleBourbon.Size,
			&singleBourbon.Price,
			&singleBourbon.Abv,
			&singleBourbon.Description,
		)

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

func GetBourbonById(id string) (Bourbon, error) {
	qs := `SELECT id, name, price, size, abv, description
			FROM bourbons
			WHERE id = ?`

	stmt, err := DB.Prepare(qs)

	if err != nil {
		return Bourbon{}, err
	}

	b := Bourbon{}

	stmtErr := stmt.QueryRow(id).Scan(&b.Id, &b.Name, &b.Price, &b.Size, &b.Abv, &b.Description)
	if stmtErr != nil {
		if stmtErr == sql.ErrNoRows {
			return Bourbon{}, nil
		}
		return Bourbon{}, stmtErr
	}

	return b, nil
}

func CreateBourbon(nb Bourbon) (bool, error) {
	tx, err := DB.Begin()

	if err != nil {
		return false, err
	}

	qs := `INSERT INTO 
			bourbons (name, price, size, abv, description) 
			VALUES (?, ?, ?, ?, ?)`

	stmt, err := tx.Prepare(qs)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(nb.Name, nb.Price, nb.Size, nb.Abv, nb.Description)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func UpdateBourbon(id int, ub Bourbon) (bool, error) {

	_, dbErr := DB.Begin()

	if dbErr != nil {
		return false, dbErr
	}

	conv := BourbonToMap(ub)

	qs := sq.Update("bourbons").Where(sq.Eq{"id": id})

	query := qs.SetMap(conv)

	result, err := query.RunWith(DB).Exec()
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()

	if count == 0 {
		return false, err

	}
	return true, nil

}

func DeleteBourbon(id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	qs := `DELETE from bourbons WHERE id = ?`

	stmt, err := DB.Prepare(qs)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func BourbonToMap(b Bourbon) map[string]interface{} {
	var output = map[string]interface{}{}
	data, _ := json.Marshal(b)
	json.Unmarshal(data, &output)

	return output

}
