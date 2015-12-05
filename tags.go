package mandrill

import (
	"encoding/json"
)

type Tags struct {
	m *Mandrill
}

// List returns all of the user-defined tag information
func (t *Tags) List() ([]tagInfo, error) {
	var ret []tagInfo
	body, err := t.m.execute("/tags/list.json", simpleRequest{t.m.APIKey})
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

type tagInfo struct {
	// the actual tag as a string
	Tag string `json:"tag"`

	// the tag's current reputation on a scale from 0 to 100
	Reputation int `json:"reputation"`

	// the total number of messages sent with this tag
	Sent int `json:"sent"`

	// the total number of hard bounces by messages with this tag
	HardBounces int `json:"hard_bounces"`

	// the total number of soft bounces by messages with this tag
	SoftBounces int `json:"soft_bounces"`

	// the total number of rejected messages with this tag
	Rejects int `json:"rejects"`

	// the total number of spam complaints received for messages with this tag
	Complaints int `json:"complaints"`

	// the total number of unsubscribe requests received for messages with this tag
	Unsubs int `json:"unsubs"`

	// the total number of times messages with this tag have been opened
	Opens int `json:"opens"`

	// the total number of times tracked URLs in messages with this tag have been clicked
	Clicks int `json:"clicks"`

	// the number of unique opens for emails sent with this tag
	UniqueOpens int `json:"unique_opens"`

	// the number of unique clicks for emails sent with this tag
	UniqueClicks int `json:"unique_clicks"`
}

// Delete permanently removes a tag, its stats and from any messages that have been sent
// There is no way to undo this operation, so use it carefully.
func (t *Tags) Delete(tag string) (tagInfo, error) {
	var ret tagInfo
	body, err := t.m.execute("/tags/delete.json", struct {
		APIKey string `json:"key"`
		Tag    string `json:"tag"`
	}{t.m.APIKey, tag})
	if err != nil {
		return ret, err
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

// Info returns more detailed information about a single tag, including aggregates of recent stats
func (t *Tags) Info(tag string) (tagStats, error) {
	var ret tagStats
	body, err := t.m.execute("/tags/info.json", struct {
		APIKey string `json:"key"`
		Tag    string `json:"tag"`
	}{t.m.APIKey, tag})
	if err != nil {
		return ret, err
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

// tagStats is very similar to tagInfo but we replicate a lot of the field
// for clarity, that way every field that is accessible is known to have a
// value (i.e. no nil pointers or empty fields)
type tagStats struct {
	Tag          string  `json:"tag"`
	Reputation   int     `json:"reputation"`
	Sent         int     `json:"sent"`
	HardBounces  int     `json:"hard_bounces"`
	SoftBounces  int     `json:"soft_bounces"`
	Rejects      int     `json:"rejects"`
	Complaints   int     `json:"complaints"`
	Unsubs       int     `json:"unsubs"`
	Opens        int     `json:"opens"`
	Clicks       int     `json:"clicks"`
	UniqueOpens  int     `json:"unique_opens"`
	UniqueClicks int     `json:"unique_clicks"`
	Stats        tagStat `json:"stat"`
}

type tagStat struct {
	Today      tagStatDetail `json:"today"`
	Last7Days  tagStatDetail `json:"last_7_days"`
	Last30Days tagStatDetail `json:"last_30_days"`
	Last60Days tagStatDetail `json:"last_60_days"`
	Last90Days tagStatDetail `json:"last_90_days"`
}

type tagStatDetail struct {
	Reputation   int `json:"reputation"`
	Sent         int `json:"sent"`
	HardBounces  int `json:"hard_bounces"`
	SoftBounces  int `json:"soft_bounces"`
	Rejects      int `json:"rejects"`
	Complaints   int `json:"complaints"`
	Unsubs       int `json:"unsubs"`
	Opens        int `json:"opens"`
	Clicks       int `json:"clicks"`
	UniqueOpens  int `json:"unique_opens"`
	UniqueClicks int `json:"unique_clicks"`
}

// TimeSeries returns hourly stats for the last 30 days for a tag
func (t *Tags) TimeSeries(tag string) ([]tagTimeSeries, error) {
	var ret []tagTimeSeries
	body, err := t.m.execute("/tags/time-series.json", struct {
		APIKey string `json:"key"`
		Tag    string `json:"tag"`
	}{t.m.APIKey, tag})
	if err != nil {
		return ret, err
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

type tagTimeSeries struct {
	Time         string `json:"time"`
	Sent         int    `json:"sent"`
	HardBounces  int    `json:"hard_bounces"`
	SoftBounces  int    `json:"soft_bounces"`
	Rejects      int    `json:"rejects"`
	Complaints   int    `json:"complaints"`
	Unsubs       int    `json:"unsubs"`
	Opens        int    `json:"opens"`
	UniqueOpens  int    `json:"unique_opens"`
	Clicks       int    `json:"clicks"`
	UniqueClicks int    `json:"unique_clicks"`
}

// AllTimeSeries returns hourly stats for the last 30 days for all tags
func (t *Tags) AllTimeSeries() ([]tagTimeSeries, error) {
	var ret []tagTimeSeries
	body, err := t.m.execute("/tags/all-time-series.json", simpleRequest{t.m.APIKey})
	if err != nil {
		return ret, err
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}
