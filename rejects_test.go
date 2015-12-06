package mandrill

import (
	"testing"
)

func TestRejects(t *testing.T) {
	// add email to blacklist
	m := NewMandrill(TestAPIKey)
	e := "reject@test.mandrillapp.com"
	if resp, err := m.Rejects().Add(e, "test reject", ""); err != nil {
		t.Error(err)
		return
	} else if resp.Email != e || !resp.Added {
		t.Errorf("failed to add blacklisted email. Response: %s", resp)
		return
	}

	// list blacklisted emails
	if resp, err := m.Rejects().List("", true, ""); err != nil {
		t.Error(err)
		return
	} else if len(resp) == 0 {
		t.Error("failed to retrieve any blacklisted email")
		return
	}

	// delete blacklisted email
	if resp, err := m.Rejects().Delete(e, ""); err != nil {
		t.Error(err)
		return
	} else if resp.Email != e || !resp.Deleted {
		t.Errorf("failed to delete blacklisted email. Response: %s", resp)
		return
	}
}
