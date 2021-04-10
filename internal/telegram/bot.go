package telegram

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

type Bot struct {
	Client *Client
}

func (b *Bot) StartListen() {
	server := &http.Server{Addr: ":3000", Handler: http.HandlerFunc(b.Client.AddHandler)}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	err := b.stopListen(server)
	if err != nil {
		fmt.Println(err)
	}
	wg.Wait()
}

func (b *Bot) stopListen(server *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func InitBot(client *Client) *Bot {
	return &Bot{Client: client}
}
