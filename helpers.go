package apiai

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) httpCall(method string, url string, body io.Reader) (
	data []byte, err error) {

	// construct request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		err = fmt.Errorf("couldn't instantiate new API.ai request with URL %v: %v", url, err)
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	client := &http.Client{}

	// perform request
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("couldn't perform API.ai query request %+v: %+v", req, err)
		return
	}
	defer resp.Body.Close()

	// parse response body
	if resp.Body != nil {
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			err = fmt.Errorf("couldn't read API.ai response body: %+v", err)
			return
		}
		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
			err = fmt.Errorf("unsuccessful response status %d from API.ai: %v",
				resp.StatusCode, string(data))
			return
		}
	}
	return
}

func httpBody(val string) io.Reader {
	return strings.NewReader(val)
}
