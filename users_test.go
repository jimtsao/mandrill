package mandrill

import (
	"testing"
)

func TestGetInfo(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	info, err := m.Users().Info()
	if err != nil {
		t.Errorf("%s. Received: %+v\n", err, info)
		return
	}
}

func TestPing(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	alive, err := m.Users().Ping()
	if err != nil {
		t.Error(err)
	} else if !alive {
		t.Error("failed to receive pong")
	}
}

func TestSenders(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	s, err := m.Users().Senders()
	if err != nil {
		t.Errorf("error: %s. data: %+v", err, s)
	}
}
