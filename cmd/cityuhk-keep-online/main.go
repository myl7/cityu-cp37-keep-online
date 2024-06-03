package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	cityuhk_ko "github.com/myl7/cityuhk-keep-online"
)

func main() {
	_ = godotenv.Load()

	err := cityuhk_ko.Login(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
}
