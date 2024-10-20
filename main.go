package main

import (
	"fmt"
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

var listMenu = []Food{
	{
		name:  "nasi goreng",
		price: 15.000,
		tipe:  FoodType,
	},
	{
		name:  "orange juice",
		price: 5.000,
		tipe:  DrinkType,
	},
	{
		name:  "lele goreng",
		price: 13.000,
		tipe:  FoodType,
	},
}

type User struct {
	username string
	password string
	status   int
}

type TypeMenu string

const (
	FoodType  TypeMenu = "food"
	DrinkType TypeMenu = "drink"
)

type Food struct {
	name  string
	price float64
	tipe  TypeMenu
}

func main() {
	menu() // Memanggil menu utama
}

func menu() {
	var nomor int
	var user User
	fmt.Println("Silahkan Pilih Nomor Dari Pilihan Berikut :\n" +
		"1. Login \n" +
		"2. Register")
	fmt.Print("Masukkan Nomor :")
	fmt.Scanln(&nomor)

	switch nomor {
	case 1:
		user.login(&listUser) // Panggil login
		break
	case 2:
		register(&listUser) // Panggil register
		break
	default:
		fmt.Println("Anda Memasukkan Nomor yang Salah")
		break
	}
}

func (user User) login(users *[]User) {
	fmt.Print("Masukkan Username = ")
	fmt.Scanln(&user.username)
	fmt.Print("Masukkan Password = ")
	fmt.Scanln(&user.password)

	switch user.checkUser(users) {
	case 1:
		Admin()
		break
	case 2:
		fmt.Println("Ini Adalah Client")
		break
	default:
		fmt.Println("Username atau Password yang Anda Masukkan Salah")
		break
	}
}

func (user User) checkUser(users *[]User) int {
	for _, item := range *users {
		if user.username == item.username && user.password == item.password {
			return item.status
		}
	}
	return 0
}

func register(users *[]User) {
	Adduser(users) // Tambahkan user baru
	menu()         // Kembali ke menu setelah register
}

func Adduser(users *[]User) {
	var user User
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&user.username)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&user.password)
	fmt.Println("Silahkan Pilih Role: \n1. Admin\n2. Client\nMasukkan nomor role:")
	fmt.Scanln(&user.status)

	*users = append(*users, user) // Tambahkan user baru ke slice

	fmt.Println("User Baru Terdaftar:", *users)
}

func Admin() {
	var nomor int
	fmt.Println(
		"___________________________________________________\n" +
			"|-----------------Restoran Apa Saja----------------| \n" +
			"|--------------------------------------------------|\n" +
			"| 1.List Menu______________________________________|\n" +
			"| 2.Add Menu_______________________________________|\n" +
			"| 3.List User______________________________________|\n" +
			"| 4.List Discount__________________________________|\n" +
			"|__________________________________________________|\n" +
			"Silahkan Pilih Nomor = ")
	fmt.Scanln(&nomor)
}
