package redirect

import (
	"net/http"

	"github.com/xDarkicex/PortfolioGo/app/controllers/filter"
	"github.com/xDarkicex/PortfolioGo/helpers"
)

//HTTPS will redirect https traffic too http
func HTTPS(w http.ResponseWriter, r *http.Request) {
	remoteIP := helpers.GetIP(r)
	err := filter.IP(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	target := "https://rolofson.me" + r.URL.EscapedPath()
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	helpers.SilentLogger.Printf("remoteIP: %s\ntarget: %s\n", remoteIP, target)
	http.Redirect(w, r, target, http.StatusMovedPermanently)
}
