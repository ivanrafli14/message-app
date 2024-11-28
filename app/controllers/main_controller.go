package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ivanrafli14/fast-campus/message-app/pkg/env"
)

func RenderHello(c *fiber.Ctx) error {
	awsPublicIp := env.GetEnv("AWS_PUBLIC_IP", "127.0.0.1")

	return c.Render("index", fiber.Map{
		"awsPublicIP": awsPublicIp,
	})
}

func RenderRegister(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}
