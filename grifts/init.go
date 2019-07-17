package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/sadegh-shahbazi/vzk/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
