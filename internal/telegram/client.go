package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"telegramBot/domain"
	"telegramBot/pkg/parser"
)

type Client struct {
	BotUrl string
}

func (c Client) GetUpdates(offset int) ([]domain.Update, error) {
	resp, err := http.Get(c.BotUrl + GetUpdatesUrl + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var RestResponse domain.UpdateResponse
	err = json.Unmarshal(body, &RestResponse)
	if err != nil {
		return nil, err
	}
	return RestResponse.Result, nil
}

func (c Client) SendImage(chatId int, imageUrl string) error {
	photo := domain.Photo{
		ChatId: chatId,
		Photo:  imageUrl,
	}
	photoJson, err := json.Marshal(photo)
	if err != nil {
		return err
	}
	_, err = http.Post(c.BotUrl+SendPhotoUrl, "application/json", bytes.NewBuffer(photoJson))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddHandler(res http.ResponseWriter, req *http.Request) {
	body := &domain.WebhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("ошибка при декодировании сообщения", err)
		return
	}
	// отрефакторить
	if strings.Contains(body.Message.Text, "/mem") {
		wg.Add(1)
		memName := strings.TrimSpace(strings.Trim(body.Message.Text, "/mem"))
		go func(memName string, body *domain.WebhookReqBody) {
			defer wg.Done()
			url, err := parser.FindMemUrl(memName)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Идет отправка изображения")
			err = c.SendImage(body.Message.Chat.ChatId, url)
			if err != nil {
				fmt.Println(err)
				return
			}
		}(memName, body)
	}
}

func NewClient(botUrl string) *Client {
	return &Client{BotUrl: botUrl}
}
