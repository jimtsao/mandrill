package mandrill

import (
	"fmt"
	"testing"
)

func TestExportsList(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Exports().List(); err != nil {
		t.Error(err)
	}
}

func TestExportsWhitelist(t *testing.T) {
	m := NewMandrill(TestAPIKey)

	if _, err := m.Exports().Whitelist(""); err == nil {
		return
	} else if ae, ok := err.(*APIError); !ok || ae.Name == "UserError" {
		fmt.Printf("[warning] should run test again at later time: %s\n", err)
	} else {
		t.Error(err)
	}
}

func TestExportsRejects(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Exports().Rejects(""); err == nil {
		return
	} else if ae, ok := err.(*APIError); !ok || ae.Name == "UserError" {
		fmt.Printf("[warning] should run test again at later time: %s\n", err)
	} else {
		t.Error(err)
	}
}

func TestExportsActivity(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	req := &ExportActivityRequest{
		NotifyEmail: "notify@51252-testmandrill.com",
	}

	if resp, err := m.Exports().Activity(req); err == nil {
		if _, err = m.Exports().Info(resp.Id); err != nil {
			t.Errorf("failed to retrieve export info. %s", err)
			return
		}
	} else if ae, ok := err.(*APIError); !ok || ae.Name == "UserError" {
		fmt.Printf("[warning] should run test again at later time: %s\n", err)
	} else {
		t.Errorf("failed to export activity. %s", err)
		return
	}
}
