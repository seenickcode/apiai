package apiai

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Query performs a simple API.ai query
func (c *Client) Query(sessionID string, q string) (answer *QueryResponse, err error) {
	q = url.QueryEscape(q)
	url := fmt.Sprintf("%s/query?v=20150910&query=%s&lang=en&sessionId=%v", APIAIBaseURL, q, sessionID)

	data, err := c.httpCall("GET", url, httpBody(""))

	answer = &QueryResponse{}
	if err := json.Unmarshal(data, &answer); err != nil {
		return answer, err
	}
	return
}

func (c *Client) httpCall(method string, url string, body io.Reader) (data []byte, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		err = fmt.Errorf("couldn't instantiate new API.ai request with URL %v: %v", url, err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("couldn't perform API.ai query request %+v: %+v", req, err)
		return
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("couldn't read API.ai response body: %+v", err)
	}
	return
}

func httpBody(val string) io.Reader {
	return strings.NewReader(val)
}
