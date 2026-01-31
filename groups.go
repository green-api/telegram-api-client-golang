package greenapi

import "encoding/json"

type GroupsCategory struct {
	GreenAPI GreenAPIInterface
}

// ------------------------------------------------------------------ CreateGroup

type RequestCreateGroup struct {
	GroupName   string   `json:"groupName"`
	ChatIds     []string `json:"chatIds"`
	Type        *string  `json:"type,omitempty"`
	IsAnonymous *bool    `json:"isAnonymous,omitempty"`
}

type CreateGroupOption func(*RequestCreateGroup) error

// Type of the group.
// parameter groupType: "group", "supergroup"
func OptionalType(groupType string) CreateGroupOption {
	return func(r *RequestCreateGroup) error {
		r.Type = &groupType
		return nil
	}
}

// If true, the group will be anonymous.
func OptionalIsAnonymous(isAnonymous bool) CreateGroupOption {
	return func(r *RequestCreateGroup) error {
		r.IsAnonymous = &isAnonymous
		return nil
	}
}

// Creating a group chat.
//
// https://green-api.com/telegram/docs/api/groups/CreateGroup/
func (c GroupsCategory) CreateGroup(groupName string, chatIds []string, options ...CreateGroupOption) (*APIResponse, error) {
	for _, chatId := range chatIds {
		err := ValidateChatId(chatId)
		if err != nil {
			return nil, err
		}
	}

	r := &RequestCreateGroup{
		GroupName: groupName,
		ChatIds:   chatIds,
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

	return c.GreenAPI.Request("POST", "createGroup", jsonData)
}

// ------------------------------------------------------------------ UpdateGroupName

type RequestUpdateGroupName struct {
	ChatId    string `json:"chatId"`
	GroupName string `json:"groupName"`
}

// Change a group chat name.
//
// https://green-api.com/telegram/docs/api/groups/UpdateGroupName/
func (c GroupsCategory) UpdateGroupName(chatId, groupName string) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestUpdateGroupName{
		ChatId:    chatId,
		GroupName: groupName,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "updateGroupName", jsonData)
}

// ------------------------------------------------------------------ UpdateGroupSettings

type RequestUpdateGroupSettings struct {
	ChatId                              string `json:"chatId"`
	AllowParticipantsSendMessages       *bool  `json:"allowParticipantsSendMessages,omitempty"`
	AllowParticipantsSendMedia          *bool  `json:"allowParticipantsSendMedia,omitempty"`
	AllowParticipantsSendPolls          *bool  `json:"allowParticipantsSendPolls,omitempty"`
	AllowParticipantsSendOtherMessages  *bool  `json:"allowParticipantsSendOtherMessages,omitempty"`
	AllowParticipantsAddWebPagePreviews *bool  `json:"allowParticipantsAddWebPagePreviews,omitempty"`
	AllowParticipantsEditGroupSettings  *bool  `json:"allowParticipantsEditGroupSettings,omitempty"`
	AllowParticipantsAddMembers         *bool  `json:"allowParticipantsAddMembers,omitempty"`
	AllowParticipantsPinMessages        *bool  `json:"allowParticipantsPinMessages,omitempty"`
}

type UpdateGroupSettingsOption func(*RequestUpdateGroupSettings) error

func OptionalAllowParticipantsSendMessages(allow bool) UpdateGroupSettingsOption {
	return func(r *RequestUpdateGroupSettings) error {
		r.AllowParticipantsSendMessages = &allow
		return nil
	}
}

func OptionalAllowParticipantsSendMedia(allow bool) UpdateGroupSettingsOption {
	return func(r *RequestUpdateGroupSettings) error {
		r.AllowParticipantsSendMedia = &allow
		return nil
	}
}

func OptionalAllowParticipantsSendPolls(allow bool) UpdateGroupSettingsOption {
	return func(r *RequestUpdateGroupSettings) error {
		r.AllowParticipantsSendPolls = &allow
		return nil
	}
}

func OptionalAllowParticipantsSendOtherMessages(allow bool) UpdateGroupSettingsOption {
	return func(r *RequestUpdateGroupSettings) error {
		r.AllowParticipantsSendOtherMessages = &allow
		return nil
	}
}

func OptionalAllowParticipantsAddWebPagePreviews(allow bool) UpdateGroupSettingsOption {
	return func(r *RequestUpdateGroupSettings) error {
		r.AllowParticipantsAddWebPagePreviews = &allow
		return nil
	}
}

func OptionalAllowParticipantsEditGroupSettings(allow bool) UpdateGroupSettingsOption {
	return func(r *RequestUpdateGroupSettings) error {
		r.AllowParticipantsEditGroupSettings = &allow
		return nil
	}
}

func OptionalAllowParticipantsAddMembers(allow bool) UpdateGroupSettingsOption {
	return func(r *RequestUpdateGroupSettings) error {
		r.AllowParticipantsAddMembers = &allow
		return nil
	}
}

func OptionalAllowParticipantsPinMessages(allow bool) UpdateGroupSettingsOption {
	return func(r *RequestUpdateGroupSettings) error {
		r.AllowParticipantsPinMessages = &allow
		return nil
	}
}

// Update group settings.
//
// https://green-api.com/telegram/docs/api/groups/UpdateGroupSettings/
func (c GroupsCategory) UpdateGroupSettings(chatId string, options ...UpdateGroupSettingsOption) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestUpdateGroupSettings{
		ChatId: chatId,
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

	return c.GreenAPI.Request("POST", "updateGroupSettings", jsonData)
}

// ------------------------------------------------------------------ GetGroupData

type RequestGetGroupData struct {
	ChatId string `json:"chatId"`
}

// Getting a group chat data
//
// https://green-api.com/telegram/docs/api/groups/GetGroupData/
func (c GroupsCategory) GetGroupData(chatId string) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestGetGroupData{
		ChatId: chatId,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "getGroupData", jsonData)
}

// ------------------------------------------------------------------ GroupParticipant

type RequestModifyGroupParticipant struct {
	ChatId            string `json:"chatId"`
	ParticipantChatId string `json:"participantChatId"`
}

// Adding a participant to a group chat.
//
// https://green-api.com/telegram/docs/api/groups/AddGroupParticipant/
func (c GroupsCategory) AddGroupParticipant(chatId, participantChatId string) (*APIResponse, error) {
	err := ValidateChatId(chatId, participantChatId)
	if err != nil {
		return nil, err
	}

	r := &RequestModifyGroupParticipant{
		ChatId:            chatId,
		ParticipantChatId: participantChatId,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "addGroupParticipant", jsonData)
}

// Removing a participant from a group chat.
//
// https://green-api.com/telegram/docs/api/groups/RemoveGroupParticipant/
func (c GroupsCategory) RemoveGroupParticipant(chatId, participantChatId string) (*APIResponse, error) {
	err := ValidateChatId(chatId, participantChatId)
	if err != nil {
		return nil, err
	}

	r := &RequestModifyGroupParticipant{
		ChatId:            chatId,
		ParticipantChatId: participantChatId,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "removeGroupParticipant", jsonData)
}

// Setting a group chat participant as an administrator.
//
// https://green-api.com/telegram/docs/api/groups/SetGroupAdmin/
func (c GroupsCategory) SetGroupAdmin(chatId, participantChatId string) (*APIResponse, error) {
	err := ValidateChatId(chatId, participantChatId)
	if err != nil {
		return nil, err
	}

	r := &RequestModifyGroupParticipant{
		ChatId:            chatId,
		ParticipantChatId: participantChatId,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "setGroupAdmin", jsonData)
}

// Removing a participant from the group chat administration rights.
//
// https://green-api.com/telegram/docs/api/groups/RemoveAdmin/
func (c GroupsCategory) RemoveAdmin(chatId, participantChatId string) (*APIResponse, error) {
	err := ValidateChatId(chatId, participantChatId)
	if err != nil {
		return nil, err
	}

	r := &RequestModifyGroupParticipant{
		ChatId:            chatId,
		ParticipantChatId: participantChatId,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "removeAdmin", jsonData)
}

// ------------------------------------------------------------------ SetGroupPicture

type RequestSetGroupPicture struct {
	File   string `json:"file"`
	ChatId string `json:"chatId"`
}

// Setting a group picture.
//
// https://green-api.com/telegram/docs/api/groups/SetGroupPicture/
func (c GroupsCategory) SetGroupPicture(filepath, chatId string) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestSetGroupPicture{
		File:   filepath,
		ChatId: chatId,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "setGroupPicture", jsonData, WithFormData(true))
}

// ------------------------------------------------------------------ LeaveGroup

type RequestLeaveGroup struct {
	ChatId string `json:"chatId"`
}

// Leaving a group chat.
//
// https://green-api.com/telegram/docs/api/groups/LeaveGroup/
func (c GroupsCategory) LeaveGroup(chatId string) (*APIResponse, error) {
	err := ValidateChatId(chatId)
	if err != nil {
		return nil, err
	}

	r := &RequestLeaveGroup{
		ChatId: chatId,
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return c.GreenAPI.Request("POST", "leaveGroup", jsonData)
}
