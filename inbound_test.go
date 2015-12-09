package mandrill

import (
	"testing"
)

func TestInboundDomains(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Inbound().Domains(); err != nil {
		t.Error(err)
	}
}

func TestInboundDomainsLifecycle(t *testing.T) {
	// add domain
	m := NewMandrill(TestAPIKey)
	d := "inbound.test.com"
	if _, err := m.Inbound().AddDomain(d); err != nil {
		t.Errorf("failed to add inbound domain. %s", err)
		return
	}

	// check domain
	if _, err := m.Inbound().CheckDomain(d); err != nil {
		t.Errorf("failed to check inbound domain. %s", err)
		return
	}

	// delete domain
	if _, err := m.Inbound().DeleteDomain(d); err != nil {
		t.Errorf("failed to delete inbound domain. %s", err)
		return
	}
}

func TestInboundRoutesLifecycle(t *testing.T) {
	// add domain
	m := NewMandrill(TestAPIKey)
	d := "meecatdev.xyz"
	if _, err := m.Inbound().AddDomain(d); err != nil {
		t.Errorf("failed to add domain. %s", err)
		return
	}

	// add route
	var id string
	if resp, err := m.Inbound().AddRoute(d, "mailbox-*", "http://google.com"); err != nil {
		t.Errorf("failed to add inbound route. %s", err)
		return
	} else {
		id = resp.Id
	}

	// update route
	if _, err := m.Inbound().UpdateRoute(id, "", ""); err != nil {
		t.Errorf("failed to update inbound route. %s", err)
		return
	}

	// retrieve route
	if _, err := m.Inbound().Routes(d); err != nil {
		t.Errorf("failed to retrieve inbound route. %s", err)
		return
	}

	// delete route
	if _, err := m.Inbound().DeleteRoute(id); err != nil {
		t.Errorf("failed to delete inbound route. %s", err)
		return
	}

	// delete domain
	if _, err := m.Inbound().DeleteDomain(d); err != nil {
		t.Errorf("failed to delete inbound domain. %s", err)
		return
	}
}

func TestInboundSendRaw(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	raw := "From: sender@example.com\nTo: mailbox-123@inbound.example.com\nSubject: Some Subject\n\nSome content."
	if _, err := m.Inbound().SendRaw(raw, nil, "", "", ""); err != nil {
		t.Error(err)
	}
}
