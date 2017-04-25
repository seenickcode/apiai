package apiai

import (
	"encoding/json"
	"time"
)

type QueryResponse struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Lang      string    `json:"lang"`
	Result    struct {
		Source           string           `json:"source"`
		ResolvedQuery    string           `json:"resolvedQuery"`
		Action           string           `json:"action"`
		ActionIncomplete bool             `json:"actionIncomplete"`
		Parameters       *json.RawMessage `json:"parameters"`
		Contexts         []*Context       `json:"contexts"`
		Metadata         struct {
			IntentID                  string `json:"intentId"`
			WebhookUsed               string `json:"webhookUsed"`
			WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
			IntentName                string `json:"intentName"`
		} `json:"metadata"`
		Fulfillment struct {
			Speech                string                  `json:"speech"`
			RawMessages           []*json.RawMessage      `json:"messages"`
			TextMessages          []*TextMessage          `json:"-""`
			ImageMessages         []*ImageMessage         `json:"-""`
			CardMessages          []*CardMessage          `json:"-""`
			QuickRepliesMessages  []*QuickRepliesMessage  `json:"-""`
			CustomPayloadMessages []*CustomPayloadMessage `json:"-"`
		} `json:"fulfillment"`
		Score float64 `json:"score"`
	} `json:"result"`
	Status    ResponseStatus `json:"status"`
	SessionID string         `json:"sessionId"`
}

type ResponseStatus struct {
	Code      int    `json:"code"`
	ErrorType string `json:"errorType"`
}

type MessageType int

// https://docs.api.ai/docs/query#section-message-objects
const (
	TextMessageType          MessageType = 0
	CardMessageType          MessageType = 1
	QuickRepliesMessageType  MessageType = 2
	ImageMessageType         MessageType = 3
	CustomPayloadMessageType MessageType = 4
)

type GenericMessage struct {
	Type MessageType `json:"type"`
}

type TextMessage struct {
	GenericMessage
	Speech string `json:"speech"`
}

type ImageMessage struct {
	GenericMessage
	ImageURL string `json:"imageUrl"`
}

type CardMessage struct {
	GenericMessage
	Title    string        `json:"title"`
	SubTitle string        `json:"subtitle"`
	Buttons  []*CardButton `json:"buttons"`
}

type QuickRepliesMessage struct {
	GenericMessage
	Title   string   `json:"title"`
	Replies []string `json:"replies"`
}

type CardButton struct {
	Text     string `json:"text"`
	Postback string `json:"postback"`
}

type CustomPayloadMessage struct {
	GenericMessage
	Payload *json.RawMessage `json:"payload"`
}
