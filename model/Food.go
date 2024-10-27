package model

import "database/sql"

type TypeMenu string

const (
	FoodType  TypeMenu = "food"
	DrinkType TypeMenu = "drink"
)

type Food struct {
	Id    string   "json:id"
	Name  string   "json:name"
	Price float64  "json:price"
	Tipe  TypeMenu "json:type"
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
		err = result.Scan(&food.Id, &food.Name, &food.Price, &food.Tipe)
		if err != nil {
			return nil, err
		}
		foods = append(foods, food)

	}

	return foods, nil
}
