package function

import (
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmplPath string, data interface{}, statusCode int) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
