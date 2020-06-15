package main

import (
	"os"

	"github.com/yagi-eng/place-search/domain/model"
	"github.com/yagi-eng/place-search/infrastructure"

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

	db, err := infrastructure.Connect()
	defer db.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&model.User{})
	db.Debug().AutoMigrate(&model.Favorite{})
}
