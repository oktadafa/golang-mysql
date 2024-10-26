package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	id    string   "json:id"
	name  string   "json:name"
	price float32  "json:price"
	tipe  TypeMenu "json:type"
}

func main() {
	db, err := sql.Open("mysql", "daffa:okta54321@tcp(localhost:3306)/golang_mysql")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	listMenusql, err := db.Query("SELECT * FROM food")
	if err != nil {
		panic(err.Error())
	}
	defer listMenusql.Close()
	var listMenuFromDb []Food
	for listMenusql.Next() {
		var menu Food
		err = listMenusql.Scan(&menu.id, &menu.name, &menu.price, &menu.tipe)
		if err != nil {
			panic(err.Error())
		}
		listMenuFromDb = append(listMenuFromDb, menu)
	}

	fmt.Println(listMenuFromDb, "list")
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
		"	 ___________________________________________________\n" +
			"|-----------------Restoran Apa Saja----------------|\n" +
			"|--------------------------------------------------|\n" +
			"| 1.Menu___________________________________________|\n" +
			"| 3.User___________________________________________|\n" +
			"| 4.Discount_______________________________________|\n" +
			"| 5.Logout_________________________________________|\n" +
			"|__________________________________________________|\n" +
			"Silahkan Pilih Nomor = ")
	fmt.Scanln(&nomor)

	switch nomor {
	case 1:
		showMenu()
		break
	case 2:
		fmt.Println("Anda Memilih Add Menu")
		break
	case 3:
		fmt.Println("Anda Memilih List User")
		break
	case 4:
		fmt.Println("Anda Memilih List Discount")
		break
	case 5:
		fmt.Println("Anda Memilih Logout")
		break
	default:
		fmt.Println("Input Anda Masukan Tidak Valid")
		break
	}
}

func showMenu() {
	var number int
	fmt.Print(
		"----------------------------------------\n" +
			"|---------------Menu-------------------|\n" +
			"|--------------------------------------|\n" +
			"|--No--|---Nama--|---Harga---|---Jenis-|\n")
	for index, item := range listMenu {
		fmt.Printf("| %d | %s | %f | %s |\n", index+1, item.name, item.price, item.tipe)
	}
	fmt.Print("|--------------------------------------|\n")
	fmt.Print(
		"Option : \n" +
			"1. Tambah Menu\n" +
			"2. Edit Menu\n" +
			"3. Hapus Menu\n" +
			"4. Exit\n")
	fmt.Print("Silahkan Pilih Nomor : ")
	fmt.Scanln(&number)
	switch number {
	case 1:
		tambahMenu()
		break
	case 2:
		fmt.Println("Anda Memilih Edit")
		break
	case 3:
		fmt.Println("Anda Memilih Hapus")
		break
	case 4:
		Admin()
		break
	default:
		fmt.Println("masukan nomor yang valid")
		showMenu()
		break
	}
}

func tambahMenu() {
	var menu Food
	var number int
	fmt.Print("Tipe Menu : \n" +
		"1. Food\n" +
		"2. Makanan\n" +
		"Silahkan Pilih Nomor Tipe Menu : ")
	fmt.Scanln(&number)
	switch number {
	case 1:
		menu.tipe = FoodType
		break
	case 2:
		menu.tipe = DrinkType
		break
	default:
		{
			fmt.Println("masukan nomor yang valid")
			tambahMenu()
		}
	}
	fmt.Print("Masukan Nama Makanan Atau Minuman : ")
	fmt.Scanln(&menu.name)

	fmt.Print("Masukan Harga : ")
	fmt.Scanln(&menu.price)

	listMenu = append(listMenu, menu)
	showMenu()
}
