package apiai

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Query performs a simple API.ai query
func (c *Client) Query(sessionID string, q string) (
	answer *QueryResponse, err error) {

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
	return
}
