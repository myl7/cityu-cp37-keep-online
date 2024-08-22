package cp37

import (
	"context"
	"log/slog"

	"github.com/go-rod/rod"
)

type CP37 struct {
	username string
	password string
}

func NewCP37(username, password string) CP37 {
	return CP37{
		username: username,
		password: password,
	}
}

const CP37LoginUrl = "https://cp37.cs.cityu.edu.hk/cp"

func (cp CP37) MustLogin(ctx context.Context, browser *rod.Browser) {
	slog.InfoContext(ctx, "cp37 login starts")

	// `MustWaitStable` never resolves. Perhaps it is because there are some failed network requests.
	page := browser.MustPage().Context(ctx)
	page.MustNavigate(CP37LoginUrl).MustWaitIdle()
	page.MustElement("#okta-signin-username").MustInput(cp.username)
	page.MustElement("#okta-signin-password").MustInput(cp.password)
	page.MustElement("#okta-signin-submit").MustClick()
	page.MustWaitStable()
	// Double waiting stable to follow a redirect.
	page.MustWaitStable()

	slog.InfoContext(ctx, "cp37 login ends")
}
