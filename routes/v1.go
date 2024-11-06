package routes

import (
	"github.com/gofiber/fiber/v2"
	"user/app/Http/Controllers"
	"user/app/Http/Middleware"
)

func SetupV1Routes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/login", Controllers.Login)
	api.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	api.Get("/users", Middleware.AccessMiddleware("view_user"), Controllers.UserIndex)
	api.Get("/users/:id", Middleware.AccessMiddleware("view_user"), Controllers.UserDetail)

	api.Get("/users/:id/permissions", Controllers.UserPermissions)
	api.Post("/users/:id/permissions", Controllers.UserSavePermissions)
}
