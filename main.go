package main

import (
	"fmt"
	"github.com/ivanrafli14/fast-campus/message-app/bootstrap"
	"github.com/ivanrafli14/fast-campus/message-app/pkg/env"
	"log"
)

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "localhost"), env.GetEnv("APP_PORT", "8080"))))
}
