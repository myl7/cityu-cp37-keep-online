package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	cityuhk_ko "github.com/myl7/cityuhk-keep-online"
)

func main() {
	_ = godotenv.Load()
	ctlUrl := os.Getenv("ROD_CTL_URL")
	username := os.Getenv("CITYUHK_USERNAME")
	if username == "" {
		log.Fatalln("Env CITYUHK_USERNAME is required")
	}
	password := os.Getenv("CITYUHK_PASSWORD")
	if password == "" {
		log.Fatalln("Env CITYUHK_PASSWORD is required")
	}

	err := cityuhk_ko.Login(context.Background(), ctlUrl, username, password)
	if err != nil {
		log.Fatalln(err)
	}
}
