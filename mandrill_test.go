package mandrill

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"testing"
)

var TestAPIKey string

func TestMain(m *testing.M) {
	flag.StringVar(&TestAPIKey, "apikey", "", "Valid Mandrill API Key")
	flag.Parse()
	if TestAPIKey == "" {
		fmt.Fprint(os.Stdout, `Please set --apikey="XYZ" flag`)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestTLS(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	tr := &http.Transport{TLSClientConfig: &tls.Config{}}
	m.HttpClient = &http.Client{Transport: tr}
	if ok, err := m.Users().Ping(); !ok || err != nil {
		t.Errorf("could not ping server through https connection")
	}
}

func TestErrInvalidKey(t *testing.T) {
	m := NewMandrill("dummy-key")
	_, err := m.Users().Ping()
	if err != ErrInvalidKey {
		t.Errorf("expected invalid api key response. %s", err)
	}
}

func TestErrValidation(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	_, err := m.execute("/users/ping.json", struct {
		BadKey string `json:"bad-key"`
	}{"bad-key"})
	if err != ErrValidation {
		t.Errorf("expected validation error response. %s", err)
	}
}
