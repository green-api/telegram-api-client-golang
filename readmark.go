package greenapi

import "encoding/json"

type ReadMarkCategory struct {
	GreenAPI GreenAPIInterface
}

// ------------------------------------------------------------------ ReadChat

type RequestReadChat struct {
	ChatId string `json:"chatId"`
}

// Marking messages in a chat as read.
//
// https://green-api.com/telegram/docs/api/marks/ReadChat/

func (c ReadMarkCategory) ReadChat(chatId string) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestReadChat{
		ChatId: chatId,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "readChat", jsonData)
}
