package helpers

import "github.com/gorilla/sessions"

var _store *sessions.CookieStore

// Store Get the mongo store
func Store() *sessions.CookieStore {
	if _store == nil {
		_store = sessions.NewCookieStore([]byte("something-very-secret"))
	}
	return _store
}
