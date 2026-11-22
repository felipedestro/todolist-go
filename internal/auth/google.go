package auth

import (
	"github.com/felipedestro/todolist/api/internal/config"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func SetupGoogle(cfg config.Config) {
	goth.UseProviders(
		google.New(
			cfg.GoogleClientID,
			cfg.GoogleClientSecret,
			cfg.GoogleCallbackURL,
			"email",
			"profile",
		),
	)
}
