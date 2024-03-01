package handler

import (
	"html/template"
	"net/http"
)

func PersonalDataHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, errParse := template.ParseFiles("WEB/Templates/personal_data.html")
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
