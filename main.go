package main

import (
	"fmt"
	"github.com/ivanrafli14/fast-campus/message-app/bootstrap"
	"github.com/ivanrafli14/fast-campus/message-app/pkg/env"
	"log"
)

//	@title			Realtime Chat Application with Go and ELK Stack
//	@version		1.0
//	@description	This project is a Realtime Chat Application built with Golang as the backend, integrated with the ELK Stack (Elasticsearch, Logstash, Kibana) for logging, monitoring, and visualization. The chat application supports real-time messaging using WebSockets and stores logs for analytics and debugging
//	@host			localhost:8080
//	@BasePath		/

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "localhost"), env.GetEnv("APP_PORT", "8080"))))
}
