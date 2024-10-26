package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oktadafa/golang-mysql/model"
	"github.com/oktadafa/golang-mysql/server"
)

// Definisikan variabel global untuk menyimpan daftar user
var listUser = []User{
	{
		username: "Admin",
		password: "Admin",
		status:   1,
	},
	{
		username: "client",
		password: "client",
		status:   2,
	},
}

type User struct {
	username string
	password string
	status   int
}

func main() {
	data, err := server.NewServer()

	if err != nil {
		panic(err.Error())
	}
	food := model.Food{}
	foodList, err := food.ListMenu(data)
	fmt.Println(foodList)

}
