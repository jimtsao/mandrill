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

	// check if it is an error response
	var errResponse apiError
	if err = json.Unmarshal(respB, &errResponse); err == nil {
		return nil, errResponse.Error()
	}

	return respB, nil
}

type apiError struct {
	Status  string
	Code    int
	Name    string
	Message string
}

var (
	ErrInvalidKey            = errors.New("The provided API key is not a valid Mandrill API key")
	ErrValidation            = errors.New("The parameters passed to the API call are invalid or not provided when required")
	ErrGeneral               = errors.New("An unexpected error occurred processing the request. Mandrill developers will be notified.")
	ErrPayment               = errors.New("The requested feature requires payment")
	ErrSubaccount            = errors.New("The provided subaccount id does not exist")
	ErrInvalidTagName        = errors.New("The requested tag does not exist or contains invalid characters")
	ErrServiceUnavailable    = errors.New("The subsystem providing this API call is down for maintenance")
	ErrUnknownSender         = errors.New("The requested sender does not exist")
	ErrUnknownURL            = errors.New("The requested URL has not been seen in a tracked link")
	ErrUnknownTrackingDomain = errors.New("The provided tracking domain does not exist")
	ErrInvalidTemplate       = errors.New("The given template name already exists or contains invalid characters")
)

func (a *apiError) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return errors.New("not an api error")
	}

	val1, ok1 := m["status"]
	val2, ok2 := m["code"]
	val3, ok3 := m["name"]
	val4, ok4 := m["message"]
	if len(m) != 4 || !ok1 || !ok2 || !ok3 || !ok4 {
		return errors.New("not an api error")
	}

	status, ok5 := val1.(string)
	code, ok6 := val2.(float64)
	name, ok7 := val3.(string)
	message, ok8 := val4.(string)

	if !ok5 || !ok6 || !ok7 || !ok8 {
		return errors.New("not an api error")
	}

	a.Status = status
	a.Code = int(code)
	a.Name = name
	a.Message = message

	return nil
}

func (a *apiError) Error() error {
	switch a.Name {
	case "Invalid_Key":
		return ErrInvalidKey
	case "ValidationError":
		return ErrValidation
	case "GeneralError":
		return ErrGeneral
	case "PaymentRequired":
		return ErrPayment
	case "Unknown_Subaccount":
		return ErrSubaccount
	case "Invalid_Tag_Name":
		return ErrInvalidTagName
	case "ServiceUnavailable":
		return ErrServiceUnavailable
	case "Unknown_Sender":
		return ErrUnknownSender
	case "Unknown_Url":
		return ErrUnknownURL
	case "Unknown_TrackingDomain":
		return ErrUnknownTrackingDomain
	case "Invalid_Template":
		return ErrInvalidTemplate
	default:
		return fmt.Errorf("An unknown error response was received from API. %+v", a)
	}
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
