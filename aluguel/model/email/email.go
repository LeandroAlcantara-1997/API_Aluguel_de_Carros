package email

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/tls"
	"fmt"

	"github.com/LeandroAlcantara-1997/model/repository"
	gomail "gopkg.in/mail.v2"
)

func RecuperarSenha(email string) error {
	senha, err := senhaAleatoria()
	if err != nil {
		return fmt.Errorf("Erro: %v", err)
	}
	
	err = repository.UpdateSenha(email, senha)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	m := gomail.NewMessage()

	// Quem envia
	m.SetHeader("From", "example@gmail.com")

	// Destinatario
	m.SetHeader("To", email)

	// Assunto
	m.SetHeader("Subject", "Recuperacao de senha")

	// corpo da mensagem
	m.SetBody("text/plain", senha)

	// Configuracao de conta
	d := gomail.NewDialer("smtp.gmail.com", 587, "example@gmail.com", "senha")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Enviando email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return nil
}

func senhaAleatoria() (string, error) {
	b := make([]byte, 10)
	_, err := rand.Read(b)

	if err != nil {
		return "", fmt.Errorf("Erro ao gerar senha %v", err)
	}

	armored := hash(b)
	return armored, err
}

func hash(b []byte) string {
	h := sha1.New()
	h.Write(b)
	sum := h.Sum(nil)
	armored := fmt.Sprintf("%x", sum)
	return armored
}
