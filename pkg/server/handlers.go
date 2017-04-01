package server

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
)

// App is the object that the views are tied to.
type App struct{}

// Index is the / of the app
func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	title := struct {
		Title string
	}{
		Title: "TEEEEEEEEEEEEEEEEESSSSSSSTTTTTTTTT",
	}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		logrus.Error(err)
	}
	if err := t.ExecuteTemplate(w, "templates/index.html", title); err != nil {
		logrus.Errorln(err)
	}
}
