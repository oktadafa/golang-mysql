package main

import (
	"database/sql"
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
	//
	if err != nil {
		panic(err.Error())
	}
	var nomor int32
	fmt.Print(
		"Silahkan Pilih Nomor: \n" +
			"1.Daftar Menu \n" +
			"2. Tambah Menu \n" +
			"3. Hapus Menu \n" +
			"Masukan Nomor : ")
	fmt.Scanln(&nomor)
	switch nomor {
	case 1:
		listMenu(data)
		break
	default:
		break
	}
}

func listMenu(data *sql.DB) {
	food := model.Food{}
	foodList, err := food.ListMenu(data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Daftar Menu")
	for number, item := range foodList {
		fmt.Printf("%d, %s..........%f \n", number, item.Name, item.Price)
	}
}
