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
		t.Errorf("failed to add blacklisted email. Response: %+v", resp)
		return
	}

	// list blacklisted emails
	if resp, err := m.Rejects().List(e, true, ""); err != nil {
		t.Error(err)
		return
	} else if len(resp) < 1 {
		t.Error("failed to retrieve any blacklisted email")
		return
	} else if resp[0].Email != e || resp[0].Detail != "test reject" {
		t.Errorf("failed to retrieve blacklisted email. Response: %+v", resp)
		return
	}

	// delete blacklisted email
	if resp, err := m.Rejects().Delete(e, ""); err != nil {
		t.Error(err)
		return
	} else if resp.Email != e || !resp.Deleted {
		t.Errorf("failed to delete blacklisted email. Response: %+v", resp)
		return
	}
}
