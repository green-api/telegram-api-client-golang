package greenapi

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

type SendingCategory struct {
	GreenAPI GreenAPIInterface
}

// ------------------------------------------------------------------ SendMessage

type RequestSendMessage struct {
	ChatId  string `json:"chatId"`
	Message string `json:"message"`
}

// Sending a text message.
//
// https://green-api.com/telegram/docs/api/sending/SendMessage/
//
// Add optional arguments by passing these functions:
//
//	OptionalQuotedMessageId(quotedMessageId string) <- Quoted message ID. If present, the message will be sent quoting the specified chat message.
//	OptionalLinkPreview(linkPreview bool) <- The parameter includes displaying a preview and a description of the link. Enabled by default.
func (c SendingCategory) SendMessage(chatId, message string) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	err = ValidateMessageLength(message, 20000)
	if err != nil {
		return nil, err
	}

	r := &RequestSendMessage{
		ChatId:  chatId,
		Message: message,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "sendMessage", jsonData)
}

// ------------------------------------------------------------------ SendFileByUpload

type RequestSendFileByUpload struct {
	ChatId   string `json:"chatId"`
	File     string `json:"file"`
	FileName string `json:"fileName"`
	Caption  string `json:"caption,omitempty"`
}

type SendFileByUploadOption func(*RequestSendFileByUpload) error

// File caption. Caption added to video, images. The telegramimum field length is 20000 characters.
func OptionalCaptionSendUpload(caption string) SendFileByUploadOption {
	return func(r *RequestSendFileByUpload) error {
		err := ValidateMessageLength(caption, 20000)
		if err != nil {
			return err
		}
		r.Caption = caption
		return nil
	}
}

// Uploading and sending a file.
//
// https://green-api.com/telegram/docs/api/sending/SendFileByUpload/
//
// Add optional arguments by passing these functions:
//
//	OptionalCaptionSendUpload(caption string) <- File caption. Caption added to video, images. The telegramimum field length is 20000 characters.
//	OptionalQuotedMessageIdSendUpload(quotedMessageId string) <- If specified, the message will be sent quoting the specified chat message.
func (c SendingCategory) SendFileByUpload(chatId, filePath, fileName string, options ...SendFileByUploadOption) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestSendFileByUpload{
		ChatId:   chatId,
		FileName: fileName,
		File:     filePath,
	}

	for _, o := range options {
		err := o(r)
		if err != nil {
			return nil, err
		}
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "sendFileByUpload", jsonData, WithFormData(true), WithMediaHost(true))
}

// ------------------------------------------------------------------ SendFileByUrl

type RequestSendFileByUrl struct {
	ChatId   string `json:"chatId"`
	UrlFile  string `json:"urlFile"`
	FileName string `json:"fileName"`
	Caption  string `json:"caption,omitempty"`
}

type SendFileByUrlOption func(*RequestSendFileByUrl) error

// File caption. Caption added to video, images. The telegramimum field length is 20000 characters.
func OptionalCaptionSendUrl(caption string) SendFileByUrlOption {
	return func(r *RequestSendFileByUrl) error {
		err := ValidateMessageLength(caption, 20000)
		if err != nil {
			return err
		}
		r.Caption = caption
		return nil
	}
}

// Sending a file by URL.
//
// https://green-api.com/telegram/docs/api/sending/SendFileByUrl/
//
// Add optional arguments by passing these functions:
//
//	OptionalCaptionSendUrl(caption string) <- File caption. Caption added to video, images. The telegramimum field length is 20000 characters.
//	OptionalQuotedMessageIdSendUrl(quotedMessageId string) <- If specified, the message will be sent quoting the specified chat message.
func (c SendingCategory) SendFileByUrl(chatId, urlFile, fileName string, options ...SendFileByUrlOption) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	err = ValidateURL(urlFile)
	if err != nil {
		return nil, err
	}

	r := &RequestSendFileByUrl{
		ChatId:   chatId,
		UrlFile:  urlFile,
		FileName: fileName,
	}

	for _, o := range options {
		err := o(r)
		if err != nil {
			return nil, err
		}
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "sendFileByUrl", jsonData)
}

// ------------------------------------------------------------------ UploadFile

type RequestUploadFile struct {
	File     []byte `json:"file"`
	FileName string `json:"fileName"`
}

// Uploading a file to the cloud storage.
//
// https://green-api.com/telegram/docs/api/sending/UploadFile/
func (c SendingCategory) UploadFile(filePath string) (*APIResponse, error) {

	binary, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "uploadFile", binary, WithSetMimetype(mtype{
		Mimetype: mimetype.Detect(binary).String(),
		FileName: filepath.Base(filePath),
	}), WithMediaHost(true))
}

// ------------------------------------------------------------------ SendPoll

type PollOption struct {
	OptionName string `json:"optionName"`
}

type RequestSendPoll struct {
	ChatId          string       `json:"chatId"`
	Message         string       `json:"message"`
	PollOptions     []PollOption `json:"options"`
	MultipleAnswers *bool        `json:"multipleAnswers,omitempty"`
}

type SendPollOption func(*RequestSendPoll) error

func OptionalMultipleAnswers(multipleAnswers bool) SendPollOption {
	return func(r *RequestSendPoll) error {
		r.MultipleAnswers = &multipleAnswers
		return nil
	}
}

// Sending messages with a poll.
//
// https://green-api.com/en/docs/api/sending/SendPoll/

func (c SendingCategory) SendPoll(chatId, message string, pollOptions []string, options ...SendPollOption) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	err = ValidateMessageLength(message, 255)
	if err != nil {
		return nil, err
	}

	if len(pollOptions) < 2 {
		return nil, fmt.Errorf("cannot create less than 2 poll options")
	} else if len(pollOptions) > 12 {
		return nil, fmt.Errorf("cannot create more than 12 poll options")
	}

	seen := make(map[string]bool)

	for _, pollOption := range pollOptions {
		if len(pollOption) > 100 {
			return nil, fmt.Errorf("poll option should not exceed 100 characters")
		}
		if seen[pollOption] {
			return nil, fmt.Errorf("poll options cannot have duplicates: %s", pollOption)
		}
		seen[pollOption] = true
	}

	r := &RequestSendPoll{
		ChatId:  chatId,
		Message: message,
	}

	for _, v := range pollOptions {
		r.PollOptions = append(r.PollOptions, PollOption{OptionName: v})
	}

	for _, o := range options {
		err := o(r)
		if err != nil {
			return nil, err
		}
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "sendPoll", jsonData)
}

// ------------------------------------------------------------------ SendLocation

type RequestSendLocation struct {
	ChatId    string  `json:"chatId"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// Sending a location message.
//
// https://green-api.com/en/docs/api/sending/SendLocation/

func (c SendingCategory) SendLocation(chatId string, latitude, longitude float32) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestSendLocation{
		ChatId:    chatId,
		Latitude:  latitude,
		Longitude: longitude,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "sendLocation", jsonData)
}

// ------------------------------------------------------------------ SendContact

type Contact struct {
	PhoneContact int    `json:"phoneContact"`
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	Company      string `json:"company,omitempty"`
}

type RequestSendContact struct {
	ChatId  string  `json:"chatId"`
	Contact Contact `json:"contact"`
}

// Sending a contact message.
//
// https://green-api.com/en/docs/api/sending/SendContact/

func (c SendingCategory) SendContact(chatId string, contact Contact) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestSendContact{
		ChatId:  chatId,
		Contact: contact,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "sendContact", jsonData)
}
