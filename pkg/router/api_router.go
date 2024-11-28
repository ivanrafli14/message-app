package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/ivanrafli14/fast-campus/message-app/app/controllers"
	"go.elastic.co/apm/module/apmfiber"
)

type ApiRouter struct {
}

func (h ApiRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api", limiter.New())
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from api",
		})
	})

	userGroup := app.Group("/user")
	userGroup.Use(apmfiber.Middleware())
	userV1Group := userGroup.Group("/v1")
	userV1Group.Post("/register", controllers.Register)
	userV1Group.Post("/login", controllers.Login)
	userV1Group.Delete("/logout", MiddlewareValidateAuth, controllers.Logout)
	userV1Group.Put("/refresh-token", MiddlewareRefreshToken, controllers.RefreshToken)

	messageGroup := app.Group("/message")
	messageGroup.Use(apmfiber.Middleware())
	messageV1Group := messageGroup.Group("/v1")
	messageV1Group.Get("/history", MiddlewareValidateAuth, controllers.GetHistory)
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}