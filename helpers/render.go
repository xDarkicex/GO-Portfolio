package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/xDarkicex/PortfolioGo/config"
)

// Render renders views blaim pug Not Secure
func Render(w http.ResponseWriter, view string) {
	if config.ENV == "development" {
		compiled, err := exec.Command("bash", "render.sh", view).Output()
		if err != nil {
			fmt.Fprintf(w, "Error: %s\n", err)
		} else {
			fmt.Fprintf(w, "%s", compiled)
			fmt.Println(view)
		}
	} else {
		ioutil.ReadFile(view)
	}
}

//RenderDynamic Dicks
func RenderDynamic(w http.ResponseWriter, view string, object interface{}) {
	if config.ENV == "development" {
		// Turn object into a json
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(object)
		compiled, err := exec.Command("bash", "render.sh", view, buf.String()).Output()
		if err != nil {
			fmt.Fprintf(w, "Error: %s\n%s", err, compiled)
		} else {
			fmt.Fprintf(w, "%s", compiled)
			fmt.Println(view)
		}
	} else {
		ioutil.ReadFile(view)
	}
}
