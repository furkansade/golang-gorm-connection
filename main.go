package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// db.Create(&User{Username: "admin2", Password: "admin"})
	// db.AutoMigrate(&User{})

	var user User
	db.First(&user, 2)
	// db.First(&user, "1")
	// db.First(&user, "password = ?", "sifre123")
	fmt.Println(user.Username)

	// var users []User
	// db.Find(&users, "password = ?", "sifre123")
	// // fmt.Println(users)
	// for _, v := range users {
	// 	fmt.Println(v)
	// }

	// var user User
	// db.First(&user, 1) // user icinde id'si 1 olani attik.
	// db.Model(&user).Update("username", "gorm")

	// var user User
	// db.First(&user, 2)
	// db.Model(&user).Updates(User{
	// 	Username: "goGorm",
	// 	Password: "gorm123",
	// })

	// var user User
	// db.Delete(&user, 3)

}
