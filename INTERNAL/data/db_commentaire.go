package data

import (
	"fmt"
	"gorm.io/gorm"
)

type Commentaire struct {
	gorm.Model
	Article string `gorm:"type:varchar(100);not null" json:"article"`
	Texte   string `gorm:"type:text;not null" json:"texte"`
}

func AjouterCommentaire(commentaire Commentaire) (Commentaire, error) {
	result := Db.Create(&commentaire) // Utilisez `Create` pour insérer une nouvelle entrée dans la base de données.
	if result.Error != nil {
		fmt.Printf("Erreur lors de l'ajout du commentaire : %v\n", result.Error)
		return Commentaire{}, result.Error
	}
	return commentaire, nil
}

func RécupérerCommentaires() ([]Commentaire, error) {
	var commentaires []Commentaire
	result := Db.Find(&commentaires) // Utilise `Find` pour récupérer tous les enregistrements.
	if result.Error != nil {
		fmt.Printf("Erreur lors de la récupération des commentaires : %v\n", result.Error)
		return nil, result.Error
	}
	return commentaires, nil
}
