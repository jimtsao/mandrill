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
	m.Metadata().Delete(name)
	if _, err := m.Metadata().Add(name, ""); err != nil {
		t.Errorf("failed to add metadata. %s", err)
		return
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
