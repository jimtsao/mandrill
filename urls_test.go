package mandrill

import (
	"testing"
)

func TestURLsList(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.URLs().List(); err != nil {
		t.Error(err)
	}
}

func TestURLsSearch(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.URLs().Search("http://test-example.com/example"); err != nil {
		t.Error(err)
	}
}

func TestURLsTrackingDomains(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.URLs().TrackingDomains(); err != nil {
		t.Error(err)
	}
}

func TestURLsAddTrackingDomain(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.URLs().AddTrackingDomain("example.com"); err != nil {
		t.Error(err)
	}
}

func TestURLsDeleteTrackingDomain(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.URLs().CheckTrackingDomain("example.com"); err != nil {
		t.Error(err)
	}
}
