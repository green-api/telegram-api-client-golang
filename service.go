package greenapi

import (
	"encoding/json"
	"fmt"
)

type ServiceCategory struct {
	GreenAPI GreenAPIInterface
}

// ------------------------------------------------------------------ CheckAccount

type RequestCheckAccount struct {
	PhoneNumber int `json:"phoneNumber"`
}

// Checking a Telegram account availability on a phone number.
//
// https://green-api.com/telegram/docs/api/service/CheckAccount/
func (c ServiceCategory) CheckAccount(phoneNumber int) (*APIResponse, error) {
	r := &RequestCheckAccount{
		PhoneNumber: phoneNumber,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "CheckAccount", jsonData)
}

// ------------------------------------------------------------------ GetAvatar

type RequestGetAvatar struct {
	ChatId string `json:"chatId"`
}

// Getting a user or a group chat avatar.
//
// https://green-api.com/telegram/docs/api/service/GetAvatar/
func (c ServiceCategory) GetAvatar(chatId string) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestGetAvatar{
		ChatId: chatId,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "getAvatar", jsonData)
}

// ------------------------------------------------------------------ GetContacts

// Getting a list of the current account contacts.
//
// https://green-api.com/telegram/docs/api/service/GetContacts/
func (c ServiceCategory) GetContacts() (*APIResponse, error) {
	return c.GreenAPI.Request("GET", "getContacts", nil)
}

// ------------------------------------------------------------------ GetContactInfo

type RequestGetContactInfo struct {
	ChatId string `json:"chatId"`
}

// Getting information about a contact.
//
// https://green-api.com/telegram/docs/api/service/GetContactInfo/
func (c ServiceCategory) GetContactInfo(chatId string) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestGetContactInfo{
		ChatId: chatId,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "getContactInfo", jsonData)
}

// ------------------------------------------------------------------ GetChats

// Getting a list of the current account chats.
//
// https://green-api.com/telegram/docs/api/service/GetChats/
func (c ServiceCategory) GetChats() (*APIResponse, error) {
	return c.GreenAPI.Request("GET", "GetChats", nil)
}

// ------------------------------------------------------------------ EditMessage

type RequestEditMessage struct {
	ChatId    string `json:"chatId"`
	Message   string `json:"message"`
	IdMessage string `json:"idMessage"`
}

// Editing a sent message.
//
// https://green-api.com/telegram/docs/api/service/EditMessage/
func (c ServiceCategory) EditMessage(chatId, idMessage, message string) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestEditMessage{
		ChatId:    chatId,
		Message:   message,
		IdMessage: idMessage,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "editMessage", jsonData)
}

// ------------------------------------------------------------------ DeleteMessage

type RequestDeleteMessage struct {
	ChatId           string `json:"chatId"`
	IdMessage        string `json:"idMessage"`
	OnlySenderDelete bool   `json:"onlySenderDelete,omitempty"`
}

// Deleting a sent message.
//
// https://green-api.com/telegram/docs/api/service/DeleteMessage/
func (c ServiceCategory) DeleteMessage(chatId, idMessage string) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestDeleteMessage{
		ChatId:    chatId,
		IdMessage: idMessage,
		// OnlySenderDelete по умолчанию false, что соответствует стандартному поведению
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "deleteMessage", jsonData)
}

// ------------------------------------------------------------------ ArchiveChat

type RequestArchiveChat struct {
	ChatId string `json:"chatId"`
}

// Archiving a chat.
//
// https://green-api.com/en/docs/api/service/archiveChat/
func (c ServiceCategory) ArchiveChat(chatId string) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestArchiveChat{
		ChatId: chatId,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "archiveChat", jsonData)
}

// Unarchiving a chat.
//
// https://green-api.com/en/docs/api/service/unarchiveChat/
func (c ServiceCategory) UnarchiveChat(chatId string) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestArchiveChat{
		ChatId: chatId,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "unarchiveChat", jsonData)
}

// ------------------------------------------------------------------ SendTyping

type RequestSendTyping struct {
	ChatId     string `json:"chatId"`
	TypingTime int    `json:"typingTime"`
	TypingType string `json:"typingType,omitempty"`
}

type SendTypingOption func(*RequestSendTyping) error

func OptionalSendTypingTime(typingTime int) SendTypingOption {
	return func(r *RequestSendTyping) error {
		if typingTime < 1000 || typingTime > 20000 {
			return fmt.Errorf("typingTime must be between 1000 and 20000 milliseconds, got: %d", typingTime)
		}
		r.TypingTime = typingTime
		return nil
	}
}

// Sending a typing status.
//
// https://green-api.com/telegram/docs/api/service/SendTyping/
func (c ServiceCategory) SendTyping(chatId string, typingTime int, options ...SendTypingOption) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestSendTyping{
		ChatId:     chatId,
		TypingTime: typingTime,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "sendTyping", jsonData)
}

// Type of typing indication.

func OptionalSendTypingType(typingType string) SendTypingOption {
	return func(r *RequestSendTyping) error {
		allowedTypes := map[string]bool{
			"text":              true,
			"record_voice_note": true,
			"upload_voice_note": true,
			"record_video_note": true,
			"upload_video_note": true,
			"record_video":      true,
			"upload_video":      true,
			"upload_photo":      true,
			"upload_document":   true,
			"choose_sticker":    true,
			"choose_location":   true,
			"choose_contact":    true,
		}
		if !allowedTypes[typingType] {
			return fmt.Errorf("invalid typingType: %s", typingType)
		}
		r.TypingType = typingType
		return nil
	}
}
