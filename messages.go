package mandrill

import (
	"encoding/json"
)

func (c *Client) MessagesSend(data *SendMessage) (responses []*Response, err error) {
	data.Key = c.Key
	body, err := c.sendApiRequest(data, "messages/send.json")
	if err != nil {
		return responses, err
	}
	responses = make([]*Response, 0)
	err = json.Unmarshal(body, &responses)
	return responses, err
}

func (c *Client) MessageSendTemplate(data *SendMessageTemplate) (responses []*Response, err error) {
	data.Key = c.Key
	body, err := c.sendApiRequest(data, "messages/send-template.json")
	if err != nil {
		return responses, err
	}
	responses = make([]*Response, 0)
	err = json.Unmarshal(body, &responses)
	return responses, err
}
