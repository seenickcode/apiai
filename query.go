package apiai

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Query performs a simple API.ai query
func (c *Client) Query(sessionID string, q string) (answer *QueryResponse, err error) {

	q = url.QueryEscape(q)
	url := fmt.Sprintf("%s/query?v=%s&query=%s&lang=en&sessionId=%v",
		APIAIBaseURL, APIVersion, q, sessionID)

	data, err := c.httpCall("GET", url, httpBody(""))
	if err != nil {
		return
	}

	answer = &QueryResponse{}
	if err := json.Unmarshal(data, &answer); err != nil {
		return answer, err
	}

	// hydrate specific fulfillment message types
	f := answer.Result.Fulfillment
	f.TextMessages = []*TextMessage{}
	f.ImageMessages = []*ImageMessage{}
	f.CardMessages = []*CardMessage{}
	f.QuickRepliesMessages = []*QuickRepliesMessage{}
	f.CustomPayloadMessages = []*CustomPayloadMessage{}
	for _, rawMsg := range f.RawMessages {

		// unmarshal to a generic message
		genMsg := &GenericMessage{}
		if err := json.Unmarshal(*rawMsg, genMsg); err != nil {
			return nil, err
		}

		// based on type value, unmarshal accordingly
		if genMsg.Type == TextMessageType {
			obj := &TextMessage{}
			if err := json.Unmarshal(*rawMsg, obj); err != nil {
				return nil, err
			}
			answer.Result.Fulfillment.TextMessages = append(answer.Result.Fulfillment.TextMessages, obj)
		} else if genMsg.Type == ImageMessageType {
			obj := &ImageMessage{}
			if err := json.Unmarshal(*rawMsg, obj); err != nil {
				return nil, err
			}
			answer.Result.Fulfillment.ImageMessages = append(answer.Result.Fulfillment.ImageMessages, obj)
		} else if genMsg.Type == CardMessageType {
			obj := &CardMessage{}
			if err := json.Unmarshal(*rawMsg, obj); err != nil {
				return nil, err
			}
			answer.Result.Fulfillment.CardMessages = append(answer.Result.Fulfillment.CardMessages, obj)
		} else if genMsg.Type == QuickRepliesMessageType {
			obj := &QuickRepliesMessage{}
			if err := json.Unmarshal(*rawMsg, obj); err != nil {
				return nil, err
			}
			answer.Result.Fulfillment.QuickRepliesMessages = append(answer.Result.Fulfillment.QuickRepliesMessages, obj)
		} else if genMsg.Type == CustomPayloadMessageType {
			obj := &CustomPayloadMessage{}
			if err := json.Unmarshal(*rawMsg, obj); err != nil {
				return nil, err
			}
			answer.Result.Fulfillment.CustomPayloadMessages = append(answer.Result.Fulfillment.CustomPayloadMessages, obj)
		}
	}
	return
}
