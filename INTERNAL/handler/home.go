package handler

import (
	"blog/INTERNAL/data"
	"fmt"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	commentaires, errData := data.RécupérerCommentaires()
	if errData != nil {
		// Si une erreur se produit, envoie une réponse HTTP avec un code d'erreur 500.
		http.Error(w, "Erreur lors de la récupération des commentaires", http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case http.MethodGet:
		tmpl, errParse := template.ParseFiles("WEB/Templates/index.html")
		if errParse != nil {
			http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
			return
		}
		errExec := tmpl.Execute(w, commentaires)
		if errExec != nil {
			http.Error(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		// Parse le formulaire.
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Erreur lors du parsing du formulaire", http.StatusBadRequest)
			return
		}

		// Récupère l'article et le texte du commentaire depuis le formulaire.
		article := r.FormValue("article")
		texte := r.FormValue("commentaire")

		// Crée une nouvelle instance de Commentaire.
		nouveauCommentaire := data.Commentaire{Article: article, Texte: texte}

		// Tente d'ajouter le nouveau commentaire dans la base de données.
		_, err := data.AjouterCommentaire(nouveauCommentaire)
		if err != nil {
			// Si une erreur se produit, envoie une réponse HTTP avec un code d'erreur 500.
			http.Error(w, fmt.Sprintf("Erreur lors de l'ajout du commentaire : %v", err), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
