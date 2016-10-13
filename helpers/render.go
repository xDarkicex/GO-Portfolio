package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	"gopkg.in/mgo.v2/bson"

	"github.com/xDarkicex/PortfolioGo/config"
	"github.com/xDarkicex/PortfolioGo/db"
	// "github.com/xDarkicex/PortfolioGo/helpers"
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

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Admin    bool          `bson:"admin"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
	// Address  Address
}

//RenderDynamic This is fancy dynamic rendering and json encoding
func RenderDynamic(req *http.Request, res http.ResponseWriter, view string, object map[string]interface{}) {
	if config.ENV == "development" {
		session, err := Store().Get(req, "user-session")
		if err != nil {
			// helpers.Logger.Println(err)
			http.Redirect(res, req, "/", 302)
			return
		}
		user := &User{}
		if session.Values["UserID"] != nil {
			err = db.Session().DB(config.ENV).C("User").FindId(bson.ObjectIdHex(session.Values["UserID"].(string))).One(&user)
			object["user"] = user
		}
		if session.Values["flash"] != nil {
			object["flash"] = session.Values["flash"].(string)
		}
		object["view"] = view
		// Turn object into a json
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(object)
		compiled, err := exec.Command("bash", "render.sh", view, buf.String()).CombinedOutput()
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
