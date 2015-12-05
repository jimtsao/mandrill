package mandrill

import (
	"encoding/json"
)

type Users struct {
	m *Mandrill
}

func (u *Users) Info() (InfoResponse, error) {
	var ret InfoResponse
	body, err := u.m.execute("/users/info.json", simpleRequest{u.m.APIKey})
	if err != nil {
		return ret, err
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

// InfoResponse contains user information including username, key, reputation, quota, and historical sending stats
type InfoResponse struct {
	// the username of the user (used for SMTP authentication)
	Username string `json:"username"`

	// the date and time that the user's Mandrill account was created as a UTC string in YYYY-MM-DD HH:MM:SS format
	CreatedAt string `json:"created_at"`

	// a unique, permanent identifier for this user
	PublicId string `json:"public_id"`

	// the reputation of the user on a scale from 0 to 100, with 75 generally being a "good" reputation
	Reputation int `json:"reputation"`

	// the maximum number of emails Mandrill will deliver for this user each hour.
	// any emails beyond that will be accepted and queued for later delivery.
	// users with higher reputations will have higher hourly quotas
	HourlyQuota int `json:"hourly_quota"`

	// the number of emails that are queued for delivery due to exceeding your monthly or hourly quotas
	Backlog int `json:"backlog"`

	// an aggregate summary of the account's sending stats
	Stats UserStats `json:"stats"`
}

type UserStats struct {
	Today      UserStat `json:"today"`
	Last7Days  UserStat `json:"last_7_days"`
	Last30Days UserStat `json:"last_30_days"`
	Last60Days UserStat `json:"last_60_days"`
	Last90Days UserStat `json:"last_90_days"`
	AllTime    UserStat `json:"all_time"`
}

type UserStat struct {
	// the number of emails sent for this user so far today
	Sent int `json:"sent"`

	// the number of emails hard bounced for this user so far today
	HardBounces int `json:"hard_bounces"`

	// the number of emails soft bounced for this user so far today
	SoftBounces int `json:"soft_bounces"`

	// the number of emails rejected for sending this user so far today
	Rejects int `json:"rejects"`

	// the number of spam complaints for this user so far today
	Complaints int `json:"complaints"`

	// the number of unsubscribes for this user so far today
	Unsubs int `json:"unsubs"`

	// the number of times emails have been opened for this user so far today
	Opens int `json:"opens"`

	// the number of unique opens for emails sent for this user so far today
	UniqueOpens int `json:"unique_opens"`

	// the number of URLs that have been clicked for this user so far today
	Clicks int `json:"clicks"`

	// the number of unique clicks for emails sent for this user so far today
	UniqueClicks int `json:"unique_clicks"`
}

// Ping checks valid connection to server with given API key
func (u *Users) Ping() (bool, error) {
	resp, err := u.m.execute("/users/ping.json", simpleRequest{u.m.APIKey})
	return string(resp) == `"PONG!"`, err
}

// Return the senders that have tried to use this account, both verified and unverified
func (u *Users) Senders() ([]Sender, error) {
	body, err := u.m.execute("/users/senders.json", simpleRequest{u.m.APIKey})
	if err != nil {
		return nil, err
	}

	var ret []Sender
	if err := json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}

	return ret, err
}

type Sender struct {
	// the sender's email address
	Address string `json:"address"`

	// the date and time that the sender was first seen by Mandrill as a UTC date string in YYYY-MM-DD HH:MM:SS format
	CreatedAt string `json:"created_at"`

	// the total number of messages sent by this sender
	Sent int `json:"sent"`

	// the total number of hard bounces by messages by this sender
	HardBounces int `json:"hard_bounces"`

	// the total number of soft bounces by messages by this sender
	SoftBounces int `json:"soft_bounces"`

	// the total number of rejected messages by this sender
	Rejects int `json:"rejects"`

	// the total number of spam complaints received for messages by this sender
	Complaints int `json:"complaints"`

	// the total number of unsubscribe requests received for messages by this sender
	Unsubs int `json:"unsubs"`

	// the total number of times messages by this sender have been opened
	Opens int `json:"opens"`

	// the total number of times tracked URLs in messages by this sender have been clicked
	Clicks int `json:"clicks"`

	// the number of unique opens for emails sent for this sender
	UniqueOpens int `json:"unique_opens"`

	// the number of unique clicks for emails sent for this sender
	UniqueClicks int `json:"unique_clicks"`
}
