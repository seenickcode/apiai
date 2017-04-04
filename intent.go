package apiai

import (
	"encoding/json"
	"fmt"
)

type Intent struct {
	ID         string             `json:"id"`
	Name       string             `json:"name"`
	ContextIn  []*json.RawMessage `json:"contextIn"`
	ContextOut []*json.RawMessage `json:"contextOut"`
	Actions    []string           `json:"actions"`
}

// TODO add remaining fields https://docs.api.ai/docs/intents#get-intents

// ListIntents fetches all intents
func (c *Client) ListIntents() (list []*Intent, err error) {

	url := fmt.Sprintf("%s/intents?v=%s", APIAIBaseURL, APIVersion)

	data, err := c.httpCall("GET", url, httpBody(""))
	if err != nil {
		return
	}

	list = []*Intent{}
	if err := json.Unmarshal(data, &list); err != nil {
		return list, err
	}
	return
}
