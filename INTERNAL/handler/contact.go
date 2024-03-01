package handler

import (
	"blog/INTERNAL/data"
	email2 "blog/INTERNAL/email"
	"fmt"
	"html/template"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, errParse := template.ParseFiles("WEB/Templates/contact.html")
		if errParse != nil {
			http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
			return
		}
		errExec := tmpl.Execute(w, nil)
		if errExec != nil {
			http.Error(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		newContact := data.Contact{Nom: name, Email: email, Text: message}
		_, err := data.AjouterContact(newContact)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors de l'ajout du commentaire : %v", err), http.StatusInternalServerError)
			return
		}

		// Construire le corps du message
		body := fmt.Sprintf("Nom: %s\nEmail: %s\nMessage: %s\n", name, email, message)

		// Envoyer l'e-mail
		email2.SendMail(body)

		http.Redirect(w, r, "/contact", http.StatusFound)
	}
}
