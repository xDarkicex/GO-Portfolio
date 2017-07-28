package helpers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/Joker/jade"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
	"gopkg.in/mgo.v2/bson"
)

var t *template.Template

//Render function renders page with our data
func Render(a *Params, view string, object map[string]interface{}) {
	fmt.Println("in render --->", a.Request.Host)
	device := a.Request.UserAgent()
	expression := regexp.MustCompile("(Mobi(le|/xyz)|Tablet)")
	if !expression.MatchString(device) {
		a.Response.Header().Set("Connection", "keep-alive")
	}
	a.Response.Header().Set("Vary", "Accept-Encoding")
	a.Response.Header().Set("Cache-Control", "private, max-age=7776000")
	a.Response.Header().Set("Transfer-Encoding", "gzip, chunked")

	object["current_user"] = a.User
	object["view"] = view

	// object["flashes"] = a.Session.Flashes()
	// fmt.Println(a.Session.Flashes())

	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	layout := Get("layout", func() *CacheObject {
		fmt.Println("inside layout")
		layout, err := jade.ParseFile("./app/views/layouts/application.pug")
		if err != nil {
			panic(err)
		}
		minified, err := m.String("text/html", layout)
		if err != nil {
			panic(err)
		}
		return NewCacheObject(minified)
	})

	currentView := Get(view, func() *CacheObject {
		fmt.Println(view)
		currentView, err := jade.ParseFile("./app/views/" + view + ".pug")
		if err != nil {
			Logger.Printf("\nParseFile error: %v", err)
		}
		minifiedView, err := m.String("text/html", currentView)
		if err != nil {
			panic(err)
		}
		return NewCacheObject(minifiedView)
	})

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
	funcMap["ParseFlashes"] = func(f []interface{}) []Flash {
		var flashes []Flash
		for _, k := range f {
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

	gotpl, err := template.New("layout").Funcs(funcMap).Parse(layout.Object.(string))
	if err != nil {
		Logger.Printf("\nTemplate parse error: %v", err)
	}
	_, err = gotpl.New(view).Parse(currentView.Object.(string))
	if err != nil {
		Logger.Printf("\nIndex parse error: %v", err)
	}
	err = gotpl.Execute(a.Response, object)
	if err != nil {
		log.Printf("\nExecute error: %v", err)
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
