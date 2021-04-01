package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"telegramBot/domain"
)

func GetUpdates(url string, offset int) ([]domain.Update, error) {
	resp, err := http.Get(url + GetUpdatesUrl + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var RestResponse domain.UpdateResponse
	err = json.Unmarshal(body, &RestResponse)
	if err != nil {
		fmt.Println(err)
	}
	return RestResponse.Result, nil
}

//TODO реализовать клиент телеграмма, для исключения передачи лишних ссылок на телеграмм бота в параметрах методов

func SendImage(chatId int, imageUrl string, botUrl string) error {
	photo := domain.Photo{
		ChatId: chatId,
		Photo:  imageUrl,
	}
	photoJson, err := json.Marshal(photo)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+SendPhotoUrl, "application/json", bytes.NewBuffer(photoJson))
	if err != nil {
		return err
	}
	return nil
}
