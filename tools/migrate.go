package main

import (
	"virtual-travel/domain/model"
	"virtual-travel/infrastructure/database"

	"github.com/sirupsen/logrus"
)

func main() {
	db, err := database.Connect()
	defer db.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&model.User{})
	db.Debug().AutoMigrate(&model.Favorite{})
}
