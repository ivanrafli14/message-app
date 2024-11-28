package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/ivanrafli14/fast-campus/message-app/pkg/database"
	"github.com/ivanrafli14/fast-campus/message-app/pkg/env"
	"github.com/ivanrafli14/fast-campus/message-app/pkg/router"
	"github.com/ivanrafli14/fast-campus/message-app/web_socket"
	"go.elastic.co/apm"
	"io"
	"log"
	"os"
)

func NewApplication() *fiber.App {
	env.SetupEnvFile()
	SetupLogfile()
	apm.DefaultTracer.Service.Name = "messaging-app"

	database.SetupDatabase()
	database.SetupMongoDB()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())

	go web_socket.ServeWSMessaging(app)

	router.InstallRouter(app)

	return app
}

func SetupLogfile() {
	logFile, err := os.OpenFile("./logs/messaging_app.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
