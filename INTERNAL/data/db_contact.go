package data

import (
	"fmt"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Nom   string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:text;not null"`
	Text  string `gorm:"type:text;not null"`
}

func AjouterContact(contact Contact) (Contact, error) {
	result := Db.Create(&contact)
	if result.Error != nil {
		fmt.Printf("Erreur lors de l'ajout du commentaire : %v\n", result.Error)
		return Contact{}, result.Error
	}
	return Contact{}, nil
}
