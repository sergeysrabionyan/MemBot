package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"telegramBot/domain"
	"telegramBot/pkg/config"
)

func getUpdates(url string, offset int) ([]domain.Update, error) {
	resp, err := http.Get(url + config.GetUpdatesUrl + "?offset=" + strconv.Itoa(offset))
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
