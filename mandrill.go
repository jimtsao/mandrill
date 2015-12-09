package mandrill

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const APIBaseURL = "https://mandrillapp.com/api/1.0/"

type Mandrill struct {
	APIKey     string
	HttpClient *http.Client
}

func NewMandrill(apikey string) Mandrill {
	return Mandrill{APIKey: apikey}
}

// simpleRequest represents requests that only require an api key
type simpleRequest struct {
	APIkey string `json:"key"`
}

// execute sends POST request to the api server
func (m *Mandrill) execute(path string, obj interface{}) ([]byte, error) {
	if obj == nil {
		return nil, errors.New("empty request")
	}
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	_, err = buf.Write(jsonBytes)
	if err != nil {
		return nil, err
	}

	url := APIBaseURL + path
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", "Mandrill Go")

	httpClient := m.HttpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respB []byte
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		g, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		respB, err = ioutil.ReadAll(g)
		if err != nil {
			return nil, err
		}
	default:
		respB, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	}

	// any non 200 is error
	if resp.StatusCode != http.StatusOK {
		var errResponse *APIError
		if err = json.Unmarshal(respB, &errResponse); err != nil {
			return nil, fmt.Errorf("failed to interpret api error. Error Response: %s", err)
		} else {
			return nil, errResponse
		}
	}

	return respB, nil
}

// FromMandrillTime returns a time struct in UTC
func FromMandrillTime(s string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", s)
}

// ToMandrillTime converts a time struct to Mandrill specific UTC format
func ToMandrillTime(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05")
}

func (m *Mandrill) Users() *Users {
	return &Users{m}
}

func (m *Mandrill) Messages() *Messages {
	return &Messages{m}
}

func (m *Mandrill) Tags() *Tags {
	return &Tags{m}
}

func (m *Mandrill) Rejects() *Rejects {
	return &Rejects{m}
}

func (m *Mandrill) Whitelists() *Whitelists {
	return &Whitelists{m}
}

func (m *Mandrill) Senders() *Senders {
	return &Senders{m}
}

func (m *Mandrill) URLs() *URLs {
	return &URLs{m}
}

func (m *Mandrill) Templates() *Templates {
	return &Templates{m}
}

func (m *Mandrill) Webhooks() *Webhooks {
	return &Webhooks{m}
}

func (m *Mandrill) Subaccounts() *Subaccounts {
	return &Subaccounts{m}
}

func (m *Mandrill) Inbound() *Inbound {
	return &Inbound{m}
}

func (m *Mandrill) Exports() *Exports {
	return &Exports{m}
}
