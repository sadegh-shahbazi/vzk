package mailers

import (
	"log"

	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr/v2"
)

var smtp mail.Sender
var r *render.Engine

func init() {

	port := envy.Get("SMTP_PORT", "587")
	host := envy.Get("SMTP_HOST", "smtp.gmail.com")
	user := envy.Get("SMTP_USER", "poshtibani980@gmail.com")
	password := envy.Get("SMTP_PASSWORD", "sadegh69sadegh69powertypowerty6218070ed")

	var err error
	smtp, err = mail.NewSMTPSender(host, port, user, password)

	if err != nil {
		log.Fatal(err)
	}

	r = render.New(render.Options{
		HTMLLayout:   "layout.html",
		TemplatesBox: packr.New("app:mailers:templates", "../templates/mail"),
		Helpers:      render.Helpers{},
	})
}
