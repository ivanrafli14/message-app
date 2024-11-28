package web_socket

import (
	"context"
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/ivanrafli14/fast-campus/message-app/app/models"
	"github.com/ivanrafli14/fast-campus/message-app/app/repository"
	"github.com/ivanrafli14/fast-campus/message-app/pkg/env"
	"go.elastic.co/apm"
	"log"
	"time"
)

func ServeWSMessaging(app *fiber.App) {
	var clients = make(map[*websocket.Conn]bool)
	var broadcast = make(chan models.MessagePayload)

	app.Get("/message/v1/send", websocket.New(func(c *websocket.Conn) {
		defer func() {
			c.Close()
			delete(clients, c)
		}()

		clients[c] = true

		for {
			var msg models.MessagePayload
			if err := c.ReadJSON(&msg); err != nil {
				log.Println("error payload: ", err)
				break
			}

			tx := apm.DefaultTracer.StartTransaction("Send Message", "ws")
			ctx := apm.ContextWithTransaction(context.Background(), tx)

			msg.Date = time.Now()
			err := repository.InsertNewMessage(ctx, msg)
			if err != nil {
				log.Println("error inserting new message: ", err)
			}
			tx.End()
			broadcast <- msg
		}
	}))

	go func() {
		for {
			msg := <-broadcast
			for client := range clients {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Println("Failed to write json: ", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}()

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "localhost"), env.GetEnv("APP_PORT_SOCKET", "8080"))))
}
