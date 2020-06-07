package main

import (
	"os"
	"virtual-travel/domain/model"
	"virtual-travel/infrastructure/database"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("ISPRD") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			logrus.Fatalf("Error loading env: %v", err)
		}
	}

	db, err := database.Connect()
	defer db.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&model.User{})
	db.Debug().AutoMigrate(&model.Favorite{})
}
