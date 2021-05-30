package main

import (
	"login/db"
	"login/model"
)

func main() {
	db := db.Connection()
	defer db.Close()

	db.AutoMigrate(&model.Login{})
	db.AutoMigrate(&model.User{})
}
