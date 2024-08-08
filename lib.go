package cp37

import (
	"time"

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

func (cp CP37) MustLogin(browser *rod.Browser) {
	// `MustWaitStable` never resolves. Maybe because of failed network requests.
	page := browser.MustPage(CP37LoginUrl).MustWaitIdle()
	page.MustElement("#okta-signin-username").MustInput(cp.username)
	page.MustElement("#okta-signin-password").MustInput(cp.password)
	page.MustElement("#okta-signin-submit").MustClick()
	page.MustWaitStable()
	time.Sleep(10 * time.Second)
}
