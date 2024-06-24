package cp37

import (
	"context"

	"github.com/go-rod/rod"
)

func Login(ctx context.Context, ctlUrl, username, password string) error {
	browser := rod.New().Context(ctx)
	if ctlUrl != "" {
		browser = browser.ControlURL(ctlUrl)
	}
	browser = browser.MustConnect()
	defer browser.MustClose()

	// `MustWaitStable` never resolves. Maybe because of failed network requests.
	page := browser.MustPage("https://cp37.cs.cityu.edu.hk/cp").MustWaitIdle()
	page.MustElement("#okta-signin-username").MustInput(username)
	page.MustElement("#okta-signin-password").MustInput(password)
	page.MustElement("#okta-signin-submit").MustClick()
	page.MustWaitStable()

	return nil
}
