package greenapi

type QueuesCategory struct {
	GreenAPI GreenAPIInterface
}

// ------------------------------------------------------------------ ShowMessagesQueue

// Getting a list of messages in the queue to be sent.
//
// https://green-api.com/telegram/docs/api/queues/ShowMessagesQueue/
func (c QueuesCategory) ShowMessagesQueue() (*APIResponse, error) {
	return c.GreenAPI.Request("GET", "showMessagesQueue", nil)
}

// ------------------------------------------------------------------ GetMessagesCount

// Getting the count of messages in the queue to be sent.
//
// https://green-api.com/telegram/docs/api/queues/GetMessagesCount/
func (c QueuesCategory) GetMessagesCount() (*APIResponse, error) {
	return c.GreenAPI.Request("GET", "getMessagesCount", nil)
}

// ------------------------------------------------------------------ ClearMessagesQueue

// Clearing the queue of messages to be sent.
//
// https://green-api.com/telegram/docs/api/queues/ClearMessagesQueue/
func (c QueuesCategory) ClearMessagesQueue() (*APIResponse, error) {
	return c.GreenAPI.Request("GET", "clearMessagesQueue", nil)
}

// ------------------------------------------------------------------ GetWebhooksCount

// Getting the count of webhooks in the queue to be sent.
//
// https://green-api.com/telegram/docs/api/queues/GetWebhooksCount/
func (c QueuesCategory) GetWebhooksCount() (*APIResponse, error) {
	return c.GreenAPI.Request("GET", "getWebhooksCount", nil)
}

// ------------------------------------------------------------------ ClearWebhooksQueue

// Clearing the queue of webhooks to be sent.
//
// https://green-api.com/telegram/docs/api/queues/ClearWebhooksQueue/
func (c QueuesCategory) ClearWebhooksQueue() (*APIResponse, error) {
	return c.GreenAPI.Request("GET", "clearWebhooksQueue", nil)
}
