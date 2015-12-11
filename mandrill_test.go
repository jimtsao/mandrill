package mandrill

import (
	"fmt"
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

func TestReadmeExamples(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	msg := &Message{
		FromName:  "John Doe",
		FromEmail: TestFromEmail,
		To: []Recipient{
			{"jane@example.com", "Jane Doe", "to"},
			{"jessica@example.com", "Jessica Doe", "cc"},
		},
		Subject: "Confirmation of Membership",
		Text:    "Congratulations on becoming a member.",
	}
	_, err := m.Messages().Send(msg, false, "", nil)
	if err != nil {
		t.Error(err)
		return
	}
}
