package apiai

import (
	"encoding/json"
	"fmt"
)

type Intent struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	ContextIn  []string   `json:"contextIn"`
	ContextOut []*Context `json:"contextOut"`
	Actions    []string   `json:"actions"`
}

// TODO add remaining fields https://docs.api.ai/docs/intents#get-intents

// ListIntents fetches all intents
func (c *Client) ListIntents() (list []*Intent, err error) {

	url := fmt.Sprintf("%s/intents?v=%s", APIAIBaseURL, APIVersion)

	data, err := c.httpCall("GET", url, httpBody(""))
	if err != nil {
		return
	}
	//fmt.Printf("%v", string(data))
	list = []*Intent{}
	if err := json.Unmarshal(data, &list); err != nil {
		return list, err
	}
	return
}
