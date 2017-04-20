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

type AddContextResponse struct {
	Names  []string       `json:"names"`
	Status ResponseStatus `json:"status"`
}

// AddContext adds a context to the session
func (c *Client) AddContext(sessionID string, context *Context) (
	answer *AddContextResponse, err error) {

	url := fmt.Sprintf("%s/contexts?v=%s&lang=en&sessionId=%v",
		APIAIBaseURL, APIVersion, sessionID)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(context)
	data, err := c.httpCall("POST", url, b)

	answer = &AddContextResponse{}
	if err := json.Unmarshal(data, &answer); err != nil {
		return answer, err
	}
	if answer.Status.Code != http.StatusOK {
		err = fmt.Errorf("API.ai response code was %d", answer.Status.Code)
		return answer, err
	}
	return
}

// FetchContexts Retrieves the list of all currently active contexts for the specified session.
func (c *Client) FetchContexts(sessionID string) (
	contexts []*Context, err error) {

	url := fmt.Sprintf("%s/contexts?v=%s&sessionId=%v",
		APIAIBaseURL, APIVersion, sessionID)

	data, err := c.httpCall("GET", url, httpBody(""))

	contexts = []*Context{}
	if err := json.Unmarshal(data, &contexts); err != nil {
		return contexts, err
	}
	return
}

// DeleteContext Deletes the specified context from the specified session.
func (c *Client) DeleteContext(sessionID string, name string) (
	status *ResponseStatus, err error) {

	url := fmt.Sprintf("%s/contexts/%s?v=%s&sessionId=%v",
		APIAIBaseURL, name, APIVersion, sessionID)

	data, err := c.httpCall("DELETE", url, httpBody(""))

	status = &ResponseStatus{}
	if err := json.Unmarshal(data, &status); err != nil {
		return status, err
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
