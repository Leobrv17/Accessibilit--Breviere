package email

import (
	"fmt"
	"net/smtp"
)

func SendMail(body string) {
	from := "leo.breviere@cgchallenge.fr"
	password := "Ja1melesp1zza!"

	to := "leo.breviere@ynov.com"
	subject := "Requête bien reçu"

	// Configuration du serveur SMTP
	smtpServer := "ssl0.ovh.net"
	smtpPort := "587"

	// Authentification de l'expéditeur
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// Construire le message à envoyer
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Envoi du message
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		fmt.Println("Erreur lors de l'envoi du message:", err)
		return
	}
	fmt.Println("Message envoyé avec succès !")
}
