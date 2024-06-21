package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	cityuhk_ko "github.com/myl7/cityuhk-keep-online"
)

const timeout = 60 * time.Second

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

	ctx, concel := context.WithTimeout(context.Background(), timeout)
	defer func() {
		concel()
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatalln("Login timeouts after", timeout)
		}
	}()
	err := cityuhk_ko.Login(ctx, ctlUrl, username, password)
	if err != nil {
		log.Fatalln(err)
	}
}
