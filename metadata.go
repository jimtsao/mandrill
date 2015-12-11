package mandrill

import (
	"encoding/json"
)

type Metadata struct {
	m *Mandrill
}

func (m *Metadata) List() ([]metadataResponse, error) {
	var ret []metadataResponse
	data := simpleRequest{m.m.APIKey}
	body, err := m.m.execute("/metadata/list.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type metadataResponse struct {
	Name         string `json:"name"`
	State        string `json:"state"`
	ViewTemplate string `json:"view_template"`
}

func (m *Metadata) Add(name string, viewTemplate string) (metadataResponse, error) {
	var ret metadataResponse
	data := struct {
		APIKey       string `json:"key"`
		Name         string `json:"name"`
		ViewTemplate string `json:"view_template,omitempty"`
	}{m.m.APIKey, name, viewTemplate}
	body, err := m.m.execute("/metadata/add.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (m *Metadata) Update(name string, viewTemplate string) (metadataResponse, error) {
	var ret metadataResponse
	data := struct {
		APIKey       string `json:"key"`
		Name         string `json:"name"`
		ViewTemplate string `json:"view_template"`
	}{m.m.APIKey, name, viewTemplate}
	body, err := m.m.execute("/metadata/update.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (m *Metadata) Delete(name string) (metadataResponse, error) {
	var ret metadataResponse
	data := struct {
		APIKey string `json:"key"`
		Name   string `json:"name"`
	}{m.m.APIKey, name}
	body, err := m.m.execute("/metadata/delete.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}
