package telegram

import (
	"context"
	"log"
	"net/http"
)

type Bot struct {
	Client *Client
}

func (b *Bot) Start(ctx context.Context) {
	go func() {
		err := http.ListenAndServe(":3000", http.HandlerFunc(b.Client.AddHandler))
		if err != nil {
			log.Fatal(err)
		}
	}()
	// Todo в планах добавить ожидание завершения всех рутин/операций
	<-ctx.Done()
}

func InitBot(client *Client) *Bot {
	return &Bot{Client: client}
}
