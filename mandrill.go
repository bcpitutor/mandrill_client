package mandrill

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (err Error) Error() string {
	return err.Message
}

func ClientWithKey(key string) *Client {
	return &Client{
		Key:        key,
		HTTPClient: &http.Client{},
		BaseURL:    "https://mandrillapp.com/api/1.0/",
	}
}

func (c *Client) sendApiRequest(data interface{}, path string) (body []byte, err error) {
	payload, _ := json.Marshal(data)
	resp, err := c.HTTPClient.Post(c.BaseURL+path, "application/json", bytes.NewReader(payload))
	if err != nil {

		return body, err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {

		return body, err
	}

	if resp.StatusCode >= 400 {
		resError := &Error{}
		err = json.Unmarshal(body, resError)
		return body, resError
	}

	return body, err
}

func (c *Client) Ping() (pong string, err error) {
	var data struct {
		Key string `json:"key"`
	}

	data.Key = c.Key

	body, err := c.sendApiRequest(data, "users/ping.json")
	if err != nil {
		return pong, err
	}

	err = json.Unmarshal(body, &pong)
	return pong, err
}
