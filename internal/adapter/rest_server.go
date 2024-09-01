package adapter

import (
	"github.com/gofiber/fiber"
	"github.com/rs/zerolog/log"
)

func WithRestServer(app *fiber.App) Option {
	log.Info().Msg("Rest server connected")
	return func(a *Adapter) {
		a.RestServer = app
	}
}
