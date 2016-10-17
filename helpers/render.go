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
func Render(a RouterArgs, view string, object map[string]interface{}) {
	if config.ENV == "development" {
		session, err := Store().Get(a.Request, "user-session")
		if err != nil {
			// helpers.Logger.Println(err)
			http.Redirect(a.Response, a.Request, "/", 302)
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
			fmt.Fprintf(a.Response, "Error: %s\n%s", err, compiled)
			Logger.Println(err)
		} else {
			fmt.Fprintf(a.Response, "%s", compiled)
			fmt.Printf("Rendering %s dynamically\n", view)
		}
	} else {
		ioutil.ReadFile(view)
	}
}

// User struct for passing a user everywhere
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Admin    bool          `bson:"admin"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

//RenderDynamic THIS IS NOW A DEPRECATED FUNC!!, Left for now for backwards protection
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
