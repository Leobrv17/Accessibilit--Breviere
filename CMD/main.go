package main

import (
	"blog/INTERNAL/data"
	"blog/INTERNAL/handler"
	"fmt"
	"net/http"
)

func main() {
	data.InitDB()
	fmt.Println("(http://localhost:8080/) - Server started on port 8080")

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/contact", handler.ContactHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/personal_data", handler.PersonalDataHandler)
	http.HandleFunc("/faq", handler.FaqHandler)

	fsStatic := http.FileServer(http.Dir("WEB/Static"))
	http.Handle("/static/", http.StripPrefix("/static/", fsStatic))
	fsStatic = http.FileServer(http.Dir("WEB/Image"))
	http.Handle("/image/", http.StripPrefix("/image/", fsStatic))
	// Démarrer le serveur
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
