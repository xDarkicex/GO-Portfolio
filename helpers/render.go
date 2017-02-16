package helpers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/Joker/jade"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
	"github.com/xDarkicex/PortfolioGo/config"
)

// var pugEngine *html.Engine
var t *template.Template
var cachedLayoutMap = make(map[int]string)

func cacheLayout() error {
	layout, err := jade.ParseFile("./app/views/layouts/application.pug")
	if err != nil {
		return err
	}
	cachedLayoutMap[1] = layout
	return nil
}

//Render function renders page with our data
func Render(a RouterArgs, view string, object map[string]interface{}) {
	times := make(map[string]interface{})
	times["total"] = time.Now()

	object["current_user"] = a.User
	object["view"] = view

	// times["gorilla-save"] = time.Now()
	object["flashes"] = a.Session.Flashes()
	a.Session.Save(a.Request, a.Response)
	// times["gorilla-save"] = time.Since(times["gorilla-save"].(time.Time))

	times["jade"] = time.Now()
	err := cacheLayout()
	if err != nil {
		panic(err)
	}
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	layoutMin, err := m.String("text/html", cachedLayoutMap[1])
	if err != nil {
		panic(err)
	}
	currentView, err := jade.ParseFile("./app/views/" + view + ".pug")
	if err != nil {
		Logger.Printf("\nParseFile error: %v", err)
	}
	currentViewMin, err := m.String("text/html", currentView)
	if err != nil {
		panic(err)
	}
	times["jade"] = time.Since(times["jade"].(time.Time))

	////////////////////////////////////////////
	// FUNCMAPS ARE LIFE!!! THIS IS LIFE NOW
	////////////////////////////////////////////

	funcMap := make(template.FuncMap)
	funcMap["Split"] = func(s string, d string) []string {
		return strings.Split(s, d)
	}
	funcMap["Join"] = func(a []string, b string) string {
		return strings.Join(a, b)
	}
	funcMap["ParseFlashes"] = func(fucks []interface{}) []Flash {
		var flashes []Flash
		for _, k := range fucks {
			var flash Flash
			json.Unmarshal([]byte(k.(string)), &flash)
			flashes = append(flashes, flash)
		}
		return flashes
	}
	funcMap["formatPostTime"] = func(t time.Time) string {
		return t.Format(time.UnixDate)
	}

	funcMap["formatTitle"] = func(s string) string {
		title := strings.SplitAfter(s, "/")
		return strings.Title(title[1])
	}

	times["render-page"] = time.Now()
	gotpl, err := template.New("layout").Funcs(funcMap).Parse(layoutMin)
	if err != nil {
		Logger.Printf("\nTemplate parse error: %v", err)
	}
	_, err = gotpl.New(view).Parse(currentViewMin)
	if err != nil {
		Logger.Printf("\nIndex parse error: %v", err)
	}
	err = gotpl.Execute(a.Response, object)
	if err != nil {
		log.Printf("\nExecute error: %v", err)
	}
	times["render-page"] = time.Since(times["render-page"].(time.Time))

	times["total"] = time.Since(times["total"].(time.Time))
	if config.Data.Env != "production" {
		fmt.Println("Render Start ==>")
		defer fmt.Println("Render Complete ==> ", times["total"])
		for k, v := range times {
			fmt.Printf("Time %s: %s\n", k, v)
		}
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
