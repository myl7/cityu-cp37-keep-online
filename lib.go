package cityuhk_ko

import (
	"context"
	"os"

	"github.com/go-rod/rod"
)

func Login(ctx context.Context) error {
	browser := rod.New().Context(ctx).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://cp37.cs.cityu.edu.hk/cp").MustWaitStable()
	page.MustElement("#okta-signin-username").MustInput(os.Getenv("CITYUHK_USERNAME"))
	page.MustElement("#okta-signin-password").MustInput(os.Getenv("CITYUHK_PASSWORD"))
	page.MustElement("#okta-signin-submit").MustClick()

	return nil
}
