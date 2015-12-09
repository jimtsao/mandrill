package mandrill

import (
	"encoding/json"
)

type Webhooks struct {
	m *Mandrill
}

type webhooksResponse struct {
	// a unique integer indentifier for the webhook
	Id int `json:"id"`

	// The URL that the event data will be posted to
	URL string `json:"url"`

	// a description of the webhook
	Description string `json:"description"`

	// the key used to requests for this webhook
	AuthKey string `json:"auth_key"`

	// The message events that will be posted to the hook
	// send, hard_bounce, soft_bounce, open, click, spam, unsub, or reject
	Events []string `json:"events"`

	// the date and time that the webhook was created as a UTC string
	CreatedAt string `json:"created_at"`

	// the date and time that the webhook last successfully received events as a UTC string
	LastSentAt string `json:"last_sent_at"`

	// the number of event batches that have ever been sent to this webhook
	BatchesSent int `json:"batches_sent"`

	// the total number of events that have ever been sent to this webhook
	EventsSent int `json:"events_sent"`

	// if we've ever gotten an error trying to post to this webhook, the last error that we've seen
	LastError string `json:"last_error"`
}

func (w *Webhooks) List() ([]webhooksResponse, error) {
	var ret []webhooksResponse
	data := simpleRequest{w.m.APIKey}
	body, err := w.m.execute("/webhooks/list.json", data)
	if err != nil {
		return ret, err
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

func (w *Webhooks) Add(url string, description string, events []string) (webhooksResponse, error) {
	var ret webhooksResponse
	data := struct {
		APIKey      string   `json:"key"`
		URL         string   `json:"url"`
		Description string   `json:"description,omitempty"`
		Events      []string `json:"events,omitempty"`
	}{w.m.APIKey, url, description, events}
	body, err := w.m.execute("/webhooks/add.json", data)
	if err != nil {
		return ret, err
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

func (w *Webhooks) Info(id int) (webhooksResponse, error) {
	var ret webhooksResponse
	data := struct {
		APIKey string `json:"key"`
		Id     int    `json:"id"`
	}{w.m.APIKey, id}
	body, err := w.m.execute("/webhooks/info.json", data)
	if err != nil {
		return ret, err
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

func (w *Webhooks) Update(id int, url string, description string, events []string) (webhooksResponse, error) {
	var ret webhooksResponse
	data := struct {
		APIKey      string   `json:"key"`
		Id          int      `json:"id"`
		URL         string   `json:"url"`
		Description string   `json:"description,omitempty"`
		Events      []string `json:"events,omitempty"`
	}{w.m.APIKey, id, url, description, events}
	body, err := w.m.execute("/webhooks/update.json", data)
	if err != nil {
		return ret, err
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

func (w *Webhooks) Delete(id int) (webhooksResponse, error) {
	var ret webhooksResponse
	data := struct {
		APIKey string `json:"key"`
		Id     int    `json:"id"`
	}{w.m.APIKey, id}
	body, err := w.m.execute("/webhooks/delete.json", data)
	if err != nil {
		return ret, err
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}
