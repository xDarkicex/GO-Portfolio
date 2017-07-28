package helpers

import (
	"bytes"
	"encoding/json"

	"github.com/gorilla/sessions"
	"github.com/xDarkicex/PortfolioGo/config"
)

var _store *sessions.CookieStore

// Store Get the session cookie
func Store() *sessions.CookieStore {
	if _store == nil {
		_store = sessions.NewCookieStore([]byte(config.Data.Secret))
	}
	return _store
}

//AddFlash Add a new flash to sessions
func AddFlash(a *Params, f Flash) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(f)
	session, _ := _store.Get(a.Request, "user-session")
	session.AddFlash(buf.String())
	// Flashes are an a array of strings.
}
