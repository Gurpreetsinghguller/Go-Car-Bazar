package Utilities

import "html/template"

func GetTemplate() *template.Template {
	tmpl := template.Must(template.ParseGlob("template/*"))
	return tmpl
}
