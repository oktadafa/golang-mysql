package model

import "database/sql"

type TypeMenu string

const (
	FoodType  TypeMenu = "food"
	DrinkType TypeMenu = "drink"
)

type Food struct {
	id    string   "json:id"
	name  string   "json:name"
	price float64  "json:price"
	tipe  TypeMenu "json:type"
}

func (food *Food) ListMenu(server *sql.DB) ([]Food, error) {
	var err error
	var foods []Food

	result, err := server.Query("SELECT * FROM food")
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var food Food
		err = result.Scan(&food.name, &food.id, &food.price, &food.tipe)
		if err != nil {
			return nil, err
		}
		foods = append(foods, food)

	}

	return foods, nil
}
