package mandrill

import (
	"encoding/json"
)

type Exports struct {
	m *Mandrill
}

func (e *Exports) Info(id string) (exportsResponse, error) {
	var ret exportsResponse
	data := struct {
		APIKey string `json:"key"`
		Id     string `json:"id"`
	}{e.m.APIKey, id}
	body, err := e.m.execute("/exports/info.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type exportsResponse struct {
	Id         string `json:"id"`
	CreatedAt  string `json:"created_at"`
	Type       string `json:"type"`
	FinishedAt string `json:"finished_at"`
	State      string `json:"state"`
	ResultURL  string `json:"result_url"`
}

func (e *Exports) List() ([]exportsResponse, error) {
	var ret []exportsResponse
	data := simpleRequest{e.m.APIKey}
	body, err := e.m.execute("/exports/list.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (e *Exports) Rejects(notifyEmail string) (exportsResponse, error) {
	var ret exportsResponse
	data := struct {
		APIKey      string `json:"key"`
		NotifyEmail string `json:"notify_email,omitempty"`
	}{e.m.APIKey, notifyEmail}
	body, err := e.m.execute("/exports/rejects.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (e *Exports) Whitelist(notifyEmail string) (exportsResponse, error) {
	var ret exportsResponse
	data := struct {
		APIKey      string `json:"key"`
		NotifyEmail string `json:"notify_email,omitempty"`
	}{e.m.APIKey, notifyEmail}
	body, err := e.m.execute("/exports/whitelist.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type ExportActivityRequest struct {
	// an optional email address to notify when the export job has finished
	NotifyEmail string `json:"notify_email,omitempty"`

	// start date as a UTC string
	DateFrom string `json:"date_from,omitempty"`

	// end date as a UTC string
	DateTo string `json:"date_to,omitempty"`

	// an array of tag names to narrow the export to; will match messages that contain ANY of the tags
	Tags []string `json:"tags,omitempty"`

	// an array of senders to narrow the export to
	Senders []string `json:"senders,omitempty"`

	// an array of states to narrow the export to; messages with ANY of the states will be included
	States []string `json:"states,omitempty"`

	// an array of api keys to narrow the export to; messsagse sent with ANY of the keys will be included
	APIKeys []string `json:"api_keys,omitempty"`
}

// Activity begins export of your activity history. The activity will be exported to a zip archive containing
// a single file named activity.csv in the same format as you would be able to export from your account's
// activity view. It includes the following fields: Date, Email Address, Sender, Subject, Status, Tags,
// Opens, Clicks, Bounce Detail. If you have configured any custom metadata fields, they will be included
// in the exported data.
func (e *Exports) Activity(request *ExportActivityRequest) (exportsResponse, error) {
	var ret exportsResponse
	type fakeRequest ExportActivityRequest
	data := struct {
		APIKey string `json:"key"`
		fakeRequest
	}{e.m.APIKey, fakeRequest(*request)}
	body, err := e.m.execute("/exports/activity.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}
