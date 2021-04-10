package telegram

import (
	"context"
	"log"
	"net/http"
	"sync"
)

var wg = &sync.WaitGroup{}

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
	<-ctx.Done()
	wg.Wait()
}

func InitBot(client *Client) *Bot {
	return &Bot{Client: client}
}
