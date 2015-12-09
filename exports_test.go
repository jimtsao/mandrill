package mandrill

import (
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
	if _, err := m.Exports().Whitelist(""); err != nil {
		t.Error(err)
	}
}

func TestExportsRejects(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Exports().Rejects(""); err != nil {
		t.Error(err)
	}
}

func TestExportsActivity(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	req := &ExportActivityRequest{
		NotifyEmail: "notify@51252-testmandrill.com",
	}
	var id string
	if resp, err := m.Exports().Activity(req); err != nil {
		t.Errorf("failed to export activity. %s", err)
		return
	} else {
		id = resp.Id
	}

	if _, err := m.Exports().Info(id); err != nil {
		t.Errorf("failed to retrieve export info. %s", err)
		return
	}
}
