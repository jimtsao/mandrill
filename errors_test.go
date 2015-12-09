package mandrill

import (
	"testing"
)

func TestErrInvalidKey(t *testing.T) {
	m := NewMandrill("dummy-key")
	_, err := m.Users().Ping()
	if err != ErrInvalidKey {
		t.Errorf("expected invalid api key error. Received: %s", err)
	}
}

func TestErrValidation(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	_, err := m.execute("/users/ping.json", struct {
		BadKey string `json:"bad-key"`
	}{"bad-key"})
	if err != ErrValidation {
		t.Errorf("expected validation error. Received: %s", err)
	}
}

func TestErrUnknownSubaccount(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	msg := &Message{
		FromEmail:  TestFromEmail,
		To:         []Recipient{{Email: "test@test.mandrillapp.com"}},
		SubAccount: "test_bad_subaccount",
	}
	if _, err := m.Messages().Send(msg, false, "", nil); err != ErrSubaccount {
		t.Error("expected unknown subaccount error. Received: %s", err)
	}
}

func TestErrInvalidTag(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Tags().Info("5219351-test-tag"); err != ErrInvalidTagName {
		t.Error("expected invalid tag error. Received: %s", err)
	}
}

func TestErrUnknownSender(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Senders().Info("bad@sender.test"); err != ErrUnknownSender {
		t.Error("expected unknown sender error. Received: %s", err)
	}
}

func TestErrUnknownURL(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.URLs().TimeSeries("bad-test-url"); err != ErrUnknownURL {
		t.Error("expected unknown url error. Received: %s", err)
	}
}

func TestErrUnknownTrackingDomain(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.URLs().CheckTrackingDomain("bad-test-domain"); err != ErrUnknownTrackingDomain {
		t.Error("expected unknown tracking domain error. Received: %s", err)
	}
}

func TestErrInvalidTemplate(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Templates().Info("521521-bad-test-template"); err != ErrUnknownTemplate {
		t.Error("expected unknown template error. Received: %s", err)
	}
}
