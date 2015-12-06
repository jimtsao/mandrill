package mandrill

import (
	"encoding/json"
)

type Senders struct {
	m *Mandrill
}

type txtRecord struct {
	// whether this domain's record is valid for use with Mandrill
	Valid bool `json:"valid"`

	// when the domain's record will be considered valid for use with Mandrill
	// as a UTC string. If set, this indicates that the record is valid now, but
	// was previously invalid, and Mandrill will wait until the record's TTL
	// elapses to start using it.
	ValidAfter string `json:"valid_after"`

	// an error describing the record, or empty if the record is correct
	Error string `json:"error"`
}

// List senders that have tried to use this account
func (s *Senders) List() ([]sendersListResponse, error) {
	var ret []sendersListResponse
	data := simpleRequest{s.m.APIKey}
	body, err := s.m.execute("/senders/list.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// sendersListResponse has data for each sending addresses used by the account
type sendersListResponse struct {
	Address      string `json:"address"`
	CreatedAt    string `json:"created_at"`
	Sent         int    `json:"sent"`
	HardBounces  int    `json:"hard_bounces"`
	SoftBounces  int    `json:"soft_bounces"`
	Rejects      int    `json:"rejects"`
	Complaints   int    `json:"complaints"`
	Unsubs       int    `json:"unsubs"`
	Opens        int    `json:"opens"`
	Clicks       int    `json:"clicks"`
	UniqueOpens  int    `json:"unique_opens"`
	UniqueClicks int    `json:"unique_clicks"`
}

// Domains return sender domains that have been added to this account.
func (s *Senders) Domains() ([]sendersDomain, error) {
	var ret []sendersDomain
	data := simpleRequest{s.m.APIKey}
	body, err := s.m.execute("/senders/domains.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// sendersDomain has data for each sending domain used by the account
type sendersDomain struct {
	// the sender domain name
	Domain    string `json:"domain"`
	CreatedAt string `json:"created_at"`

	// when the domain's DNS settings were last tested as a UTC string
	LastTestedAt string `json:"last_tested_at"`

	// details about the domain's SPF record
	SPF txtRecord `json:"spf"`

	// details about the domain's DKIM record
	DKIM txtRecord `json:"dkim"`

	// if the domain has been verified, when it occurred as a UTC string
	VerifiedAt string `json:"verified_at"`

	// whether this domain can be used to authenticate mail, either for itself or as a
	// custom signing domain. If this is false but spf and dkim are both valid, you will
	// need to verify the domain before using it to authenticate mail
	ValidSigning bool `json:"valid_signing"`
}

// AddDomain to your account. Sender domains are added automatically as you send,
// but you can use this call to add them ahead of time.
func (s *Senders) AddDomain(domain string) (sendersDomain, error) {
	var ret sendersDomain
	data := struct {
		APIKey string `json:"key"`
		Domain string `json:"domain"`
	}{s.m.APIKey, domain}
	body, err := s.m.execute("/senders/add-domain.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// Checks the SPF and DKIM settings for a domain. If you haven't already added
// this domain to your account, it will be added automatically.
func (s *Senders) CheckDomain(domain string) (sendersDomain, error) {
	var ret sendersDomain
	data := struct {
		APIKey string `json:"key"`
		Domain string `json:"domain"`
	}{s.m.APIKey, domain}
	body, err := s.m.execute("/senders/check-domain.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// Sends a verification email in order to verify ownership of a domain. Domain
// verification is a required step to confirm ownership of a domain. Once a domain
// has been verified in a Mandrill account, other accounts may not have their
// messages signed by that domain unless they also verify the domain. This prevents
//other Mandrill accounts from sending mail signed by your domain.
func (s *Senders) VerifyDomain(domain string, mailbox string) (sendersVerifyResponse, error) {
	var ret sendersVerifyResponse
	data := struct {
		APIKey  string `json:"key"`
		Domain  string `json:"domain"`
		Mailbox string `json:"mailbox"`
	}{s.m.APIKey, domain, mailbox}
	body, err := s.m.execute("/senders/verify-domain.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type sendersVerifyResponse struct {
	Status string `json:"status"`
	Domain string `json:"domain"`
	Email  string `json:"email"`
}

// Info returns detailed information about a single sender, including aggregates of recent stats
func (s *Senders) Info(address string) (sendersInfoResponse, error) {
	var ret sendersInfoResponse
	data := struct {
		APIKey  string `json:"key"`
		Address string `json:"address"`
	}{s.m.APIKey, address}
	body, err := s.m.execute("/senders/info.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type sendersInfoResponse struct {
	Address     string    `json:"address"`
	CreatedAt   string    `json:"created_at"`
	Sent        int       `json:"sent"`
	HardBounces int       `json:"hard_bounces"`
	SoftBounces int       `json:"soft_bounces"`
	Rejects     int       `json:"rejects"`
	Complaints  int       `json:"complaints"`
	Unsubs      int       `json:"unsubs"`
	Opens       int       `json:"opens"`
	Clicks      int       `json:"clicks"`
	Stats       UserStats `json:"stats"`
}

// TimeSeries return hourly stats for the last 30 days for a sender
func (s *Senders) TimeSeries(address string) ([]senderTimeSeries, error) {
	var ret []senderTimeSeries
	data := struct {
		APIKey  string `json:"key"`
		Address string `json:"address"`
	}{s.m.APIKey, address}
	body, err := s.m.execute("/senders/time-series.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type senderTimeSeries struct {
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
