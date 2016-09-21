package helpers

import (
	"fmt"
	"net/http"
	"os/exec"
	"regexp"

	"github.com/julienschmidt/httprouter"
)

// Render renders views
func Render(w http.ResponseWriter, view string) {
	compiled, err := exec.Command("bash", "render.sh", view).Output()
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
	} else {
		fmt.Fprintf(w, "%s", compiled)
		fmt.Println(view)
	}
}

// RenderKobraScript command to render kobrascript too javascript
func RenderKobraScript(w http.ResponseWriter, view string) {
	compiled, err := exec.Command(
		"kobrac",
		fmt.Sprintf("app/assets/kobrascripts/%s.ks", view)).Output()
	if err != nil {
		fmt.Fprintf(w, "KobraScript %s Error: %s\n", view, err)
	} else {
		w.Header().Set("Content-Type", "application/javascript;")
		fmt.Fprintf(w, "%s", compiled)
	}
}

// RenderScss for scss Render command for scss
//sass --scss -C --sourcemap=none --style=compressed app/assets/stylesheets/application.scss
func RenderScss(w http.ResponseWriter, view string) {
	compiled, err := exec.Command(
		"sass",
		"--scss",
		"-C",
		"--sourcemap=none",
		"--style=compressed",
		fmt.Sprintf("app/assets/stylesheets/%s.scss", view)).Output()
	if err != nil {
		fmt.Fprintf(w, "Stylesheet %s Error: %s\n", view, err)
	} else {
		// set server document type to css !important
		w.Header().Set("Content-Type", "text/css;")
		fmt.Fprintf(w, "%s", compiled)
	}
}

// HandleScssRequest captures scss cli output
func HandleScssRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	exp, err := regexp.CompilePOSIX("\\.css$")
	sheet := exp.ReplaceAllString(ps.ByName("sheet"), "")
	if err != nil {
		fmt.Fprintf(w, "505 Asset Error: %s\n", err)
	} else {
		RenderScss(w, sheet)
	}
}

//HandleKobraRequest Dicks
func HandleKobraRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	exp, err := regexp.CompilePOSIX("\\.js$")
	sheet := exp.ReplaceAllString(ps.ByName("sheet"), "")
	if err != nil {
		fmt.Fprintf(w, "505 Asset Error: %s\n", err)
	} else {
		RenderKobraScript(w, sheet)
	}
}
