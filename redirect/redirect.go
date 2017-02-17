package redirect
import "net/http"

//HTTPS will redirect https traffic too http
func HTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
}
