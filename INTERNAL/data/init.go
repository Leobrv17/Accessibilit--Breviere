package data

import (
	"fmt"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {
	connectDatabase()
	err := Db.AutoMigrate(&Commentaire{}, &Contact{})
	if err != nil {
		fmt.Printf("Erreur lors de la création des tables : %v\n", err)
	}
}
