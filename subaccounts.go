package mandrill

import (
	"encoding/json"
)

type Subaccounts struct {
	m *Mandrill
}

type subaccountsResponse struct {
	// a unique indentifier for the subaccount
	Id string `json:"id"`

	// an optional display name for the subaccount
	Name string `json:"name"`

	// an optional manual hourly quota for the subaccount. If not specified,
	// the hourly quota will be managed based on reputation
	CustomQuota int `json:"custom_quota"`

	// the current sending status of the subaccount, one of "active" or "paused"
	Status string `json:"status"`

	// the subaccount's current reputation on a scale from 0 to 100
	Reputation int `json:"reputation"`

	// the date and time that the subaccount was created as a UTC string
	CreatedAt string `json:"created_at"`

	// the date and time that the subaccount first sent as a UTC string
	FirstSentAt string `json:"first_sent_at"`

	// the number of emails the subaccount has sent so far this week
	// (weeks start on midnight Monday, UTC)
	SentWeekly int `json:"sent_weekly"`

	// the number of emails the subaccount has sent so far this month
	// (months start on midnight of the 1st, UTC)
	SentMonthly int `json:"sent_monthly"`

	// the number of emails the subaccount has sent since it was created
	SentTotal int `json:"sent_total"`
}

// List subaccounts defined for the account, optionally filtered by a prefix
// returns up to 1000 subaccounts
func (s *Subaccounts) List(q string) ([]subaccountsResponse, error) {
	var ret []subaccountsResponse
	data := struct {
		APIKey string `json:"key"`
		Q      string `json:"q,omitempty"`
	}{s.m.APIKey, q}
	body, err := s.m.execute("/subaccounts/list.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// Add a new subaccount. Id max length 255. Name max length 1024
func (s *Subaccounts) Add(id string, name string, notes string, quota int) (subaccountsResponse, error) {
	var ret subaccountsResponse
	data := struct {
		APIKey      string `json:"key"`
		Id          string `json:"id"`
		Name        string `json:"name,omitempty"`
		Notes       string `json:"notes,omitempty"`
		CustomQuota int    `json:"custom_quota,omitempty"`
	}{s.m.APIKey, id, name, notes, quota}
	body, err := s.m.execute("/subaccounts/add.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (s *Subaccounts) Info(id string) (subaccountsInfoResponse, error) {
	var ret subaccountsInfoResponse
	data := struct {
		APIKey string `json:"key"`
		Id     string `json:"id"`
	}{s.m.APIKey, id}
	body, err := s.m.execute("/subaccounts/info.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type subaccountsInfoResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Notes       string `json:"notes"`
	CustomQuota int    `json:"custom_quota"`
	Status      string `json:"status"`
	Reputation  int    `json:"reputation"`
	CreatedAt   string `json:"created_at"`
	FirstSentAt string `json:"first_sent_at"`
	SentWeekly  int    `json:"sent_weekly"`
	SentMonthly int    `json:"sent_monthly"`
	SentTotal   int    `json:"sent_total"`
	SentHourly  int    `json:"sent_hourly"`
	HourlyQuota int    `json:"hourly_quota"`
	Last30Days  struct {
		Clicks       int `json:"clicks"`
		Complaints   int `json:"complaints"`
		HardBounces  int `json:"hard_bounces"`
		Opens        int `json:"opens"`
		Rejects      int `json:"rejects"`
		Sent         int `json:"sent"`
		SoftBounces  int `json:"soft_bounces"`
		UniqueClicks int `json:"unique_clicks"`
		UniqueOpens  int `json:"unique_opens"`
		Unsubs       int `json:"unsubs"`
	} `json:"last_30_days"`
}

func (s *Subaccounts) Update(id string, name string, notes string, quota int) (subaccountsResponse, error) {
	var ret subaccountsResponse
	data := struct {
		APIKey      string `json:"key"`
		Id          string `json:"id"`
		Name        string `json:"name,omitempty"`
		Notes       string `json:"notes,omitempty"`
		CustomQuota int    `json:"custom_quota,omitempty"`
	}{s.m.APIKey, id, name, notes, quota}
	body, err := s.m.execute("/subaccounts/update.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (s *Subaccounts) Delete(id string) (subaccountsResponse, error) {
	var ret subaccountsResponse
	data := struct {
		APIKey string `json:"key"`
		Id     string `json:"id"`
	}{s.m.APIKey, id}
	body, err := s.m.execute("/subaccounts/delete.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// Pause a subaccount's sending. Any future emails delivered to this subaccount will
// be queued for a maximum of 3 days until the subaccount is resumed
func (s *Subaccounts) Pause(id string) (subaccountsResponse, error) {
	var ret subaccountsResponse
	data := struct {
		APIKey string `json:"key"`
		Id     string `json:"id"`
	}{s.m.APIKey, id}
	body, err := s.m.execute("/subaccounts/pause.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (s *Subaccounts) Resume(id string) (subaccountsResponse, error) {
	var ret subaccountsResponse
	data := struct {
		APIKey string `json:"key"`
		Id     string `json:"id"`
	}{s.m.APIKey, id}
	body, err := s.m.execute("/subaccounts/resume.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}
