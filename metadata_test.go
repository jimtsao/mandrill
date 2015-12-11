package mandrill

import (
	"testing"
)

func TestMetadataList(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Metadata().List(); err != nil {
		t.Error(err)
	}
}

func TestMetadataLifecycle(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	name := "test-metadata"
	if _, err := m.Metadata().Add(name, ""); err != nil {
		// there is delay in metadata handling, returns validation error if exists
		if ae, ok := err.(*APIError); !ok || ae.Name != "ValidationError" {
			t.Errorf("failed to add metadata. %s", err)
			return
		}
	}

	if _, err := m.Metadata().Update(name, "view template update"); err != nil {
		t.Errorf("failed to update metadata. %s", err)
		return
	}

	if _, err := m.Metadata().Delete(name); err != nil {
		t.Errorf("failed to delete metadata. %s", err)
		return
	}
}
