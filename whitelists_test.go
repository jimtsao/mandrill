package mandrill

import (
	"testing"
)

func TestWhitelists(t *testing.T) {
	// add email to whitelist
	m := NewMandrill(TestAPIKey)
	e := "whitelist@test.mandrillapp.com"
	if resp, err := m.Whitelists().Add(e, "test whitelist"); err != nil {
		t.Error(err)
		return
	} else if resp.Email != e || !resp.Added {
		t.Errorf("failed to add whitelist email. Response: %+v", resp)
		return
	}

	// list whitelisted emails
	if resp, err := m.Whitelists().List(e); err != nil {
		t.Error(err)
	} else if len(resp) < 1 {
		t.Error("failed to retrieve whitelisted email")
		return
	} else if resp[0].Email != e || resp[0].Detail != "test whitelist" {
		t.Error("failed to retrieve whitelisted email. Response: %+v", resp)
		return
	}

	// delete whitelisted email
	if resp, err := m.Whitelists().Delete(e); err != nil {
		t.Error(err)
	} else if resp.Email != e || !resp.Deleted {
		t.Errorf("failed to delete whitelisted email. Response: %+v", resp)
		return
	}
}
