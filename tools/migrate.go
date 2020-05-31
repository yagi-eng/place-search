package main

import (
	"virtual-travel/domain"
	"virtual-travel/infrastructure/database"

	"github.com/sirupsen/logrus"
)

func main() {
	db, err := database.Connect()
	defer db.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&domain.User{})
	db.Debug().AutoMigrate(&domain.Favorite{})
}
