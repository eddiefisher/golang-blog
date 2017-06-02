package utils

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
)

//RenderTemplate ...
func RenderTemplate(folder, name string) *template.Template {
	var (
		pattern   string
		filenames []string
	)

	filenames = append(filenames, "views/layouts/layout.tmpl")
	filenames = append(filenames, "views/layouts/menu.tmpl")
	filenames = append(filenames, strings.Join([]string{"views/", folder, "/", name, ".tmpl"}, ""))
	if len(filenames) == 0 {
		logrus.Errorf("html/template: pattern matches no files: %#q", pattern)
		return nil
	}
	t, _ := template.ParseFiles(filenames...)
	return t
}

//RenderErrorTemplate ...
func RenderErrorTemplate(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
	RenderTemplate("layouts", "error").ExecuteTemplate(w, "layout", data)
}
