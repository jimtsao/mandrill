package mandrill

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

	// check if error response
	var errResponse apiError
	json.Unmarshal(respB, &errResponse)

	return respB, errResponse.Error()
}

type apiError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

var (
	ErrInvalidKey = errors.New("The provided API key is not a valid Mandrill API key")
	ErrValidation = errors.New("The parameters passed to the API call are invalid or not provided when required")
	ErrGeneral    = errors.New("An unexpected error occurred processing the request. Mandrill developers will be notified.")
)

func (a *apiError) Error() error {
	switch a.Name {
	case "":
		return nil
	case "Invalid_Key":
		return ErrInvalidKey
	case "ValidationError":
		return ErrValidation
	case "GeneralError":
		return ErrGeneral
	default:
		return fmt.Errorf("An unknown error response was received from API. %+v", a)
	}
}

func (m *Mandrill) Users() *Users {
	return &Users{m}
}
