package main

import (
	"encoding/json"
)

func (c *Client) MessagesSend(message *Message) (responses []*Response, err error) {
	var data struct {
		Key     string   `json:"key"`
		Message *Message `json:"message,omitempty"`
		Async   bool     `json:"async,omitempty"`
		IPPool  string   `json:"ip_pool,omitempty"`
		SendAt  string   `json:"send_at,omitempty"`
	}
	data.Key = c.Key
	data.Message = message
	data.Async = message.Async
	data.IPPool = message.IPPool
	data.SendAt = message.SendAt
	body, err := c.sendApiRequest(data, "messages/send.json")
	if err != nil {
		return responses, err
	}
	responses = make([]*Response, 0)
	err = json.Unmarshal(body, &responses)
	return responses, err
}

func (c *Client) MessageSendTemplate(message *Message, templateName string, templateContent interface{}) (responses []*Response, err error) {
	var data struct {
		Key             string      `json:"key"`
		Message         *Message    `json:"message,omitempty"`
		Async           bool        `json:"async,omitempty"`
		IPPool          string      `json:"ip_pool,omitempty"`
		SendAt          string      `json:"send_at,omitempty"`
		TemplateName    string      `json:"template_name"`
		TemplateContent []*Variable `json:"template_content"`
	}
	data.Key = c.Key
	data.Message = message
	data.Async = message.Async
	data.IPPool = message.IPPool
	data.SendAt = message.SendAt
	data.TemplateName = templateName
	data.TemplateContent = ConvertMapToVariables(templateContent)

	body, err := c.sendApiRequest(data, "messages/send-template.json")
	if err != nil {
		return responses, err
	}
	responses = make([]*Response, 0)
	err = json.Unmarshal(body, &responses)
	return responses, err
}
