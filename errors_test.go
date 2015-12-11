package mandrill

import (
	"testing"
)

func TestErrInvalidKey(t *testing.T) {
	m := NewMandrill("dummy-key")
	_, err := m.Users().Ping()
	if ae, ok := err.(*APIError); !ok || ae.Name != "Invalid_Key" {
		t.Errorf("expected invalid api key error. Received: %s", err)
	}
}

func TestErrValidation(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	_, err := m.execute("/users/ping.json", struct {
		BadKey string `json:"bad-key"`
	}{"bad-key"})
	if ae, ok := err.(*APIError); !ok || ae.Name != "ValidationError" {
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
	_, err := m.Messages().Send(msg, false, "", nil)
	if ae, ok := err.(*APIError); !ok || ae.Name != "Unknown_Subaccount" {
		t.Errorf("expected unknown subaccount error. Received: %s", err)
	}

}

func TestErrInvalidTag(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	_, err := m.Tags().Info("5219351-test-tag")
	if ae, ok := err.(*APIError); !ok || ae.Name != "Invalid_Tag_Name" {
		t.Errorf("expected invalid tag error. Received: %s", err)
	}
}

func TestErrUnknownSender(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	_, err := m.Senders().Info("bad@sender.test")
	if ae, ok := err.(*APIError); !ok || ae.Name != "Unknown_Sender" {
		t.Errorf("expected unknown sender error. Received: %s", err)
	}
}

func TestErrUnknownURL(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	_, err := m.URLs().TimeSeries("bad-test-url")
	if ae, ok := err.(*APIError); !ok || ae.Name != "Unknown_Url" {
		t.Errorf("expected unknown url error. Received: %s", err)
	}
}

func TestErrUnknownTrackingDomain(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	_, err := m.URLs().CheckTrackingDomain("bad-test-domain")
	if ae, ok := err.(*APIError); !ok || ae.Name != "Unknown_TrackingDomain" {
		t.Errorf("expected unknown tracking domain error. Received: %s", err)
	}
}

func TestErrInvalidTemplate(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	_, err := m.Templates().Info("521521-bad-test-template")
	if ae, ok := err.(*APIError); !ok || ae.Name != "Unknown_Template" {
		t.Errorf("expected unknown template error. Received: %s", err)
	}
}

func TestErrUnknownWebhook(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	_, err := m.Webhooks().Info(5210129)
	if ae, ok := err.(*APIError); !ok || ae.Name != "Unknown_Webhook" {
		t.Errorf("expected unknown webhook error. Received: %s", err)
	}
}

func TestErrUnknownRoute(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	_, err := m.Inbound().UpdateRoute("51521-test-bad-route", "", "")
	if ae, ok := err.(*APIError); !ok || ae.Name != "Unknown_InboundRoute" {
		t.Errorf("expected unknown inbound route error. Received: %s", err)
	}
}

func TestErrUnknownPool(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	_, err := m.IPs().PoolInfo("test-bad-pool")
	if ae, ok := err.(*APIError); !ok || ae.Name != "Unknown_Pool" {
		t.Errorf("expected unknown pool error. Received: %s", err)
	}
}

func TestErrUnknownIP(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	_, err := m.IPs().Info("test-bad-ip")
	if ae, ok := err.(*APIError); !ok || ae.Name != "Unknown_IP" {
		t.Errorf("expected unknown ip error. Received: %s", err)
	}
}
