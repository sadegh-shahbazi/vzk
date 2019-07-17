package mailers

import (
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
)

func SendWelcomeEmails(email string) error {
	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = "آنتی لایسنس - Antilicense.com"
	m.From = "poshtibani980@gmail.com"
	m.To = []string{email}
	err := m.AddBody(r.HTML("welcome_email.html"), render.Data{})
	if err != nil {
		return err
	}
	return smtp.Send(m)
}

func SendRecoveryEmail(email string, recoveryString string) error {
	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = "آنتی لایسنس - بازیابی کلمه عبور - Antilicense.com"
	m.From = "poshtibani980@gmail.com"
	m.To = []string{email}

	err := m.AddBody(r.HTML("welcome_email.html"), render.Data{"recoveryString": recoveryString})
	if err != nil {
		return err
	}
	return smtp.Send(m)
}
