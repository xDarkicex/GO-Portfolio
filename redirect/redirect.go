package redirect

import (
	"log"
	"net/http"
)

//HTTPS will redirect https traffic too http
func HTTPS(w http.ResponseWriter, r *http.Request) {
	target := "https://" + r.Host + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, r, target, http.StatusPermanentRedirect)
}
