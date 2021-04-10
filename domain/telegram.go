package domain

type Update struct {
	Id      int     `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type UpdateResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

type Photo struct {
	ChatId int    `json:"chat_id"`
	Photo  string `json:"photo"`
}

type WebhookReqBody struct {
	Message Message `json:"message"`
}
