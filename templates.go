package mandrill

import (
	"encoding/json"
)

type Templates struct {
	m *Mandrill
}

type Template struct {
	Name      string   `json:"name"`
	FromEmail string   `json:"from_email,omitempty"`
	FromName  string   `json:"from_name,omitempty"`
	Subject   string   `json:"subject,omitempty"`
	Code      string   `json:"code,omitempty"`
	Text      string   `json:"text,omitempty"`
	Publish   bool     `json:"publish,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}

func (t *Templates) Add(template *Template) (templateResponse, error) {
	var ret templateResponse
	type fakeTemplate Template
	data := struct {
		APIKey string `json:"key"`
		fakeTemplate
	}{t.m.APIKey, fakeTemplate(*template)}
	body, err := t.m.execute("/templates/add.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type templateResponse struct {
	// the immutable unique code name of the template
	Slug string `json:"slug"`

	// the name of the template
	Name string `json:"name"`

	// the list of labels applied to the template
	Labels []string `json:"labels"`

	// the full HTML code of the template, with mc:edit attributes marking the
	// editable elements - draft version
	Code string `json:"code"`

	// the subject line of the template, if provided - draft version
	Subject string `json:"subject"`

	// the default sender address for the template, if provided - draft version
	FromEmail string `json:"from_email"`

	// the default sender from name for the template, if provided - draft version
	FromName string `json:"from_name"`

	// the default text part of messages sent with the template, if provided
	// - draft version
	Text string `json:"text"`

	// the same as the template name - kept as a separate field for backwards
	// compatibility
	PublishName string `json:"publish_name"`

	// the full HTML code of the template, with mc:edit attributes marking the
	// editable elements that are available as published, if it has been published
	PublishCode string `json:"publish_code"`

	// the subject line of the template, if provided
	PublishSubject string `json:"publish_subject"`

	// the default sender address for the template, if provided
	PublishFromEmail string `json:"publish_from_email"`

	// the default sender from name for the template, if provided
	PublishFromName string `json:"publish_from_name"`

	// the default text part of messages sent with the template, if provided
	PublishText string `json:"publish_text"`

	// the date and time the template was last published as a UTC string,
	// or null if it has not been published
	PublishedAt string `json:"published_at"`

	// the date and time the template was first created as a UTC string
	CreatedAt string `json:"created_at"`

	// the date and time the template was last modified as a UTC string
	UpdatedAt string `json:"updated_at"`
}

func (t *Templates) Info(name string) (templateResponse, error) {
	var ret templateResponse
	data := struct {
		APIKey string `json:"key"`
		Name   string `json:"name"`
	}{t.m.APIKey, name}
	body, err := t.m.execute("/templates/info.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (t *Templates) Update(template *Template) (templateResponse, error) {
	var ret templateResponse
	type fakeTemplate Template
	data := struct {
		APIKey string `json:"key"`
		fakeTemplate
	}{t.m.APIKey, fakeTemplate(*template)}
	body, err := t.m.execute("/templates/update.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (t *Templates) Publish(name string) (templateResponse, error) {
	var ret templateResponse
	data := struct {
		APIKey string `json:"key"`
		Name   string `json:"name"`
	}{t.m.APIKey, name}
	body, err := t.m.execute("/templates/publish.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (t *Templates) Delete(name string) (templateResponse, error) {
	var ret templateResponse
	data := struct {
		APIKey string `json:"key"`
		Name   string `json:"name"`
	}{t.m.APIKey, name}
	body, err := t.m.execute("/templates/delete.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (t *Templates) List(label string) ([]templateResponse, error) {
	var ret []templateResponse
	data := struct {
		APIKey string `json:"key"`
		Label  string `json:"label"`
	}{t.m.APIKey, label}
	body, err := t.m.execute("/templates/list.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (t *Templates) TimeSeries(name string) ([]templatesTimeSeries, error) {
	var ret []templatesTimeSeries
	data := struct {
		APIKey string `json:"key"`
		Name   string `json:"name"`
	}{t.m.APIKey, name}
	body, err := t.m.execute("/templates/time-series.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type templatesTimeSeries struct {
	Time         string `json:"time"`
	Sent         int    `json:"sent"`
	HardBounces  int    `json:"hard_bounces"`
	SoftBounces  int    `json:"soft_bounces"`
	Rejects      int    `json:"rejects"`
	Complaints   int    `json:"complaints"`
	Opens        int    `json:"opens"`
	UniqueOpens  int    `json:"unique_opens"`
	Clicks       int    `json:"clicks"`
	UniqueClicks int    `json:"unique_clicks"`
}

type TemplateMergeVar struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type TemplatesRenderRequest struct {
	TemplateName    string             `json:"template_name"`
	TemplateContent []TemplateMergeVar `json:"template_content"`
	MergeVars       []TemplateMergeVar `json:"merge_vars,omitempty"`
}

func (t *Templates) Render(r *TemplatesRenderRequest) (string, error) {
	var ret struct {
		HTML string `json:"html"`
	}
	type fakeRenderRequest TemplatesRenderRequest
	data := struct {
		APIKey string `json:"key"`
		fakeRenderRequest
	}{t.m.APIKey, fakeRenderRequest(*r)}
	body, err := t.m.execute("/templates/render.json", data)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return "", err
	}

	return ret.HTML, nil
}
