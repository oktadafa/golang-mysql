package main

import (
	"fmt"
)

type User struct {
	username string
	password string
	status   int
}

func main() {
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
	var nomor int
	var user User
	fmt.Println("Silahkan Pilih Nomor Dari Pilihan Beriku :\n" +
		"1.Login \n" +
		"2.Register")
	fmt.Print("Masukan Nomor :")
	fmt.Scanln(&nomor)
	switch nomor {
	case 1:
		user.login(listUser)
		break
	case 2:
		register(&listUser)
		break
	default:
		fmt.Println("Anda Memasukan Nomor Yang salah")
		break
	}
}

func (user User) login(users []User) {
	fmt.Print("Masukan Username = ")
	fmt.Scanln(&user.username)
	fmt.Print("Masukan Password = ")
	fmt.Scanln(&user.password)
	switch user.checkUser(users) {
	case 1:
		fmt.Println("Ini Adalah Admin")
		break
	case 2:
		fmt.Println("Ini Adalah Client")
		break
	default:
		fmt.Println("Username Atau Password Yang Masukan Salah")
		break
	}
}

func (user User) checkUser(users []User) int {

	for _, item := range users {
		if user.username == item.username && user.password == item.password {
			return item.status
		}
	}
	return 0
}
func register(users *[]User) {
	Adduser(users)
}
func Adduser(users *[]User) {
	var user User
	fmt.Print("masukan username : ")
	fmt.Scanln(&user.username)
	fmt.Print("Masukan Password : ")
	fmt.Scanln(&user.password)
	fmt.Println("Silahkan Pilih Role : \n" +
		"1.Admin\n" + "2.Client\n masukan nomor role = ")
	fmt.Scanln(&user.status)

	*users = append(*users, user)

	fmt.Println(*users)
}
