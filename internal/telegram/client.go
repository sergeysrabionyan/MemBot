package telegram

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"telegramBot/domain"
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
func NewClient(botUrl string) *Client {
	return &Client{BotUrl: botUrl}
}
