package apiai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Name       string            `json:"name"`
	Lifespan   int               `json:"lifespan"`
	Parameters map[string]string `json:"parameters"`
}

// AddContext adds a context to the session
func (c *Client) AddContext(sessionID string, context *Context) (
	answer *QueryResponse, err error) {

	url := fmt.Sprintf("%s/contexts?v=%s&lang=en&sessionId=%v",
		APIAIBaseURL, APIVersion, sessionID)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(context)
	data, err := c.httpCall("POST", url, b)

	answer = &QueryResponse{}
	if err := json.Unmarshal(data, &answer); err != nil {
		return answer, err
	}
	if answer.Status.Code != http.StatusOK {
		err = fmt.Errorf("API.ai response code was %d", answer.Status.Code)
		return answer, err
	}
	return
}

// ClearContext clears all contexts
func (c *Client) ClearContext(sessionID string) (
	answer *QueryResponse, err error) {

	url := fmt.Sprintf("%s/contexts?v=%s&lang=en&sessionId=%v",
		APIAIBaseURL, APIVersion, sessionID)

	data, err := c.httpCall("DELETE", url, nil)

	answer = &QueryResponse{}
	if err := json.Unmarshal(data, &answer); err != nil {
		return answer, err
	}
	return
}
