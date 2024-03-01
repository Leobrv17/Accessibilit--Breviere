package handler

import (
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, errParse := template.ParseFiles("WEB/Templates/about.html")
		if errParse != nil {
			http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
			return
		}
		errExec := tmpl.Execute(w, nil)
		if errExec != nil {
			http.Error(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError)
			return
		}
	}
}
