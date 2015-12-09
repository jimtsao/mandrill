package mandrill

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

var TestAPIKey string
var TestFromEmail string

func TestMain(m *testing.M) {
	TestAPIKey = os.Getenv("MANDRILL_TEST_API_KEY")
	TestFromEmail = os.Getenv("MANDRILL_TEST_FROM_EMAIL")
	if TestAPIKey == "" || TestFromEmail == "" {
		fmt.Println("Please set all ENV variables MANDRILL_TEST_API_KEY, MANDRILL_TEST_FROM_EMAIL")
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

func TestFromMandrillTime(t *testing.T) {
	t1, err := FromMandrillTime("2015-12-04 12:15:30")
	t1 = t1.UTC()
	if err != nil ||
		t1.Year() != 2015 || t1.Month() != time.December || t1.Day() != 4 ||
		t1.Hour() != 12 || t1.Minute() != 15 || t1.Second() != 30 {
		t.Errorf("expected time 4 December 2015 12:15:30, got: %s", t1)
	}
}

func TestToMandrillTime(t *testing.T) {
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		t.Error(err)
		return
	}
	t1 := ToMandrillTime(time.Date(2015, time.December, 4, 12, 15, 30, 0, loc))
	if t1 != "2015-12-04 17:15:30" {
		t.Errorf("expected time 4 December 2015 17:15:30, got: %s", t1)
	}
}
