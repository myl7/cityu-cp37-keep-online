package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/joho/godotenv"
	cp37 "github.com/myl7/cityu-cp37-keep-online"
)

func main() {
	_ = godotenv.Load()
	c := configFromEnv()

	browser := rod.New()
	defer browser.MustClose()
	if c.ctlUrl != "" {
		browser = browser.ControlURL(c.ctlUrl)
	} else {
		if c.binPath != "" {
			ctlUrl := launcher.New().Bin(c.binPath).MustLaunch()
			browser = browser.ControlURL(ctlUrl)
		} else {
			path, has := launcher.LookPath()
			if !has {
				// To work with the systemd service read-only home, we disable auto browser download
				log.Fatalln("Cannot find browser")
			}
			ctlUrl := launcher.New().Bin(path).MustLaunch()
			browser = browser.ControlURL(ctlUrl)
		}
	}
	browser = browser.MustConnect()

	ctx, concel := context.WithTimeout(context.Background(), c.timeout)
	defer func() {
		concel()
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatalf("Login timeout: timeout=%v", c.timeout)
		}
	}()

	cp37 := cp37.NewCP37(c.username, c.password)
	cp37.MustLogin(ctx, browser)
}

type config struct {
	ctlUrl   string
	binPath  string
	username string
	password string
	timeout  time.Duration
}

const TimeoutStrDefault = "30s"

func configFromEnv() config {
	ctlUrl := os.Getenv("CP37_ROD_CTL_URL")
	binPath := os.Getenv("CP37_ROD_BIN_PATH")

	username := os.Getenv("CP37_CITYU_USERNAME")
	if username == "" {
		log.Fatalln("Env CP37_CITYU_USERNAME for CityU AIMS username is required")
	}

	password := os.Getenv("CP37_CITYU_PASSWORD")
	if password == "" {
		log.Fatalln("Env CP37_CITYU_PASSWORD for CityU AIMS password is required")
	}

	timeoutStr := os.Getenv("CP37_LOGIN_TIMEOUT")
	if timeoutStr == "" {
		timeoutStr = TimeoutStrDefault
	}
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		log.Fatalf("Env CP37_LOGIN_TIMEOUT is invalid: err=%v", err)
	}
	if timeout < 0 {
		log.Fatalf("Env CP37_LOGIN_TIMEOUT is invalid: timeout=%v", timeout)
	}

	return config{
		ctlUrl:   ctlUrl,
		binPath:  binPath,
		username: username,
		password: password,
		timeout:  timeout,
	}
}
