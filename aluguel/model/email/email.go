package email

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func RecuperarSenha(email string) error {
	m := gomail.NewMessage()

	// Quem envia
	m.SetHeader("From", "emailexemplo@gmail.com")

	// Destinatario
	m.SetHeader("To", email)

	// Assunto
	m.SetHeader("Subject", "Recuperacao de senha teste")

	// corpo da mensagem
	m.SetBody("text/plain", "Recuperacao de senha")

	// Configuracao de conta
	d := gomail.NewDialer("smtp.gmail.com", 587, "emailexemplo@gmail.com", "123456")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Enviando email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return nil
}
