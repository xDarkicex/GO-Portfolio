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
			Logger.Println(err)
			fmt.Fprintf(w, "Error: %s\n", err)

		} else {
			fmt.Fprintf(w, "%s", compiled)
			fmt.Println(view)
		}
	} else {
		ioutil.ReadFile(view)
	}
}

//RenderDynamic This is fancy dynamic rendering and json encoding
func RenderDynamic(res http.ResponseWriter, view string, object interface{}) {
	if config.ENV == "development" {
		// Turn object into a json
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(object)
		compiled, err := exec.Command("bash", "render.sh", view, buf.String()).Output()
		if err != nil {
			fmt.Fprintf(res, "Error: %s\n%s", err, compiled)
			Logger.Println(err)
		} else {
			fmt.Fprintf(res, "%s", compiled)
			fmt.Printf("Rendering %s dynamically\n", view)
		}
	} else {
		ioutil.ReadFile(view)
	}
}
