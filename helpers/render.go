package helpers

import (
	"fmt"
	"net/http"
	"os/exec"
)

// Render renders views
func Render(w http.ResponseWriter, view string) {
	compiled, err := exec.Command("bash", "render.sh", view).Output()
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
	} else {
		fmt.Fprintf(w, "%s", compiled)
	}
}
