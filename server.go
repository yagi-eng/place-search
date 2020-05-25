package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env: %s", err)
	}
}

func main() {
	apiKey := os.Getenv("GMAP_API_KEY")
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.TextSearchRequest{
		Query: "東京タワー",
	}

	res, err := c.TextSearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	pretty.Println(res)
}
