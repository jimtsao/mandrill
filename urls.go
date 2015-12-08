package mandrill

import (
	"encoding/json"
)

type URLs struct {
	m *Mandrill
}

func (u *URLs) List() ([]URLsResponse, error) {
	var ret []URLsResponse
	data := simpleRequest{u.m.APIKey}
	body, err := u.m.execute("/urls/list.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type URLsResponse struct {
	// the URL to be tracked
	Clicks int `json:"clicks"`

	// the number of emails that contained the URL
	Sent int `json:"sent"`

	// the number of times the URL has been clicked from a tracked email
	UniqueClicks int `json:"unique_clicks"`

	// the number of unique emails that have generated clicks for this URL
	URL string `json:"url"`
}

func (u *URLs) Search(query string) ([]URLsResponse, error) {
	var ret []URLsResponse
	data := struct {
		APIKey string `json:"key"`
		Query  string `json:"q"`
	}{u.m.APIKey, query}
	body, err := u.m.execute("/urls/search.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (u *URLs) TimeSeries(url string) ([]URLsTimeSeriesResponse, error) {
	var ret []URLsTimeSeriesResponse
	data := struct {
		APIKey string `json:"key"`
		URL    string `json:"url"`
	}{u.m.APIKey, url}
	body, err := u.m.execute("/urls/time-series.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type URLsTimeSeriesResponse struct {
	Time         string `json:"time"`
	Sent         int    `json:"sent"`
	Clicks       int    `json:"clicks"`
	UniqueClicks int    `json:"unique_clicks"`
}

func (u *URLs) TrackingDomains() ([]URLsTrackingDomainResponse, error) {
	var ret []URLsTrackingDomainResponse
	data := simpleRequest{u.m.APIKey}
	body, err := u.m.execute("/urls/tracking-domains.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type URLsTrackingDomainResponse struct {
	Domain       string `json:"domain"`
	CreatedAt    string `json:"created_at"`
	LastTestedAt string `json:"last_tested_at"`
	Cname        struct {
		Error      string `json:"error"`
		Valid      bool   `json:"valid"`
		ValidAfter string `json:"valid_after"`
	} `json:"cname"`
	ValidTracking bool `json:"valid_tracking"`
}

func (u *URLs) AddTrackingDomain(domain string) (URLsTrackingDomainResponse, error) {
	var ret URLsTrackingDomainResponse
	data := struct {
		APIKey string `json:"key"`
		Domain string `json:"domain"`
	}{u.m.APIKey, domain}
	body, err := u.m.execute("/urls/add-tracking-domain.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (u *URLs) CheckTrackingDomain(domain string) (URLsTrackingDomainResponse, error) {
	var ret URLsTrackingDomainResponse
	data := struct {
		APIKey string `json:"key"`
		Domain string `json:"domain"`
	}{u.m.APIKey, domain}
	body, err := u.m.execute("/urls/check-tracking-domain.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}
