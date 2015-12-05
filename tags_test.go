package mandrill

import (
	"testing"
)

func TestTagsList(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Tags().List(); err != nil {
		t.Error(err)
	}
}

func TestTagsInfo(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Tags().Info("test-tag-info"); err != nil && err != ErrInvalidTagName {
		t.Error(err)
	}
}

func TestTagsTimeSeries(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Tags().TimeSeries("test-tag-time-series"); err != nil && err != ErrInvalidTagName {
		t.Error(err)
	}
}

func TestTagsTimeSeriesAll(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Tags().AllTimeSeries(); err != nil {
		t.Error(err)
	}
}

func TestTagsDelete(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Tags().Delete("test-tag-delete"); err != nil && err != ErrInvalidTagName {
		t.Error(err)
	}
}
