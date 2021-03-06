package auth

import (
	"github.com/gofiber/fiber/v2"

	"deepseen-backend/middlewares"
)

// APIs setup
func Setup(app *fiber.App) {
	group := app.Group("/api/auth")

	group.Post("/recovery/send", sendRecoveryEmail)
	group.Post("/recovery/validate", validateRecoveryCode)
	group.Get("/signout/complete", middlewares.Authorize, completeSignOut)
	group.Post("/signup", signUp)
	group.Post("/signin", signIn)
}
