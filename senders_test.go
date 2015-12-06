package mandrill

import (
	"testing"
)

func TestSendersList(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Senders().List(); err != nil {
		t.Error(err)
	}
}

func TestSendersDomains(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Senders().Domains(); err != nil {
		t.Error(err)
	}
}

func TestSendersAddDomain(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if resp, err := m.Senders().AddDomain("testdomain.test"); err != nil {
		t.Error(err)
	} else if resp.Domain != "testdomain.test" {
		t.Errorf("response domain mismatch with request. Response: %+v", resp)
	}
}

func TestSendersCheckDomain(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if resp, err := m.Senders().CheckDomain("testdomain.test"); err != nil {
		t.Error(err)
	} else if resp.Domain != "testdomain.test" {
		t.Errorf("response domain mismatch with request. Response: %+v", resp)
	}
}

func TestSendersVerifyDomain(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if resp, err := m.Senders().VerifyDomain("testdomain.test", "testuser"); err != nil {
		t.Error(err)
	} else if resp.Domain != "testdomain.test" {
		t.Errorf("response domain mismatch with request. Response: %+v", resp)
	}
}

func TestSendersInfo(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if resp, err := m.Senders().Info(TestFromEmail); err != nil && err != ErrUnknownSender {
		t.Error(err)
	} else if resp.Address != TestFromEmail {
		t.Errorf("response email address mismatch with request. Response: %+v", resp)
	}
}

func TestSendersTimeSeries(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Senders().TimeSeries(TestFromEmail); err != nil && err != ErrUnknownSender {
		t.Error(err)
	}
}
