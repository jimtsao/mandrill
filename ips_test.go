package mandrill

import (
	"testing"
)

func TestIPsList(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.IPs().List(); err != nil {
		t.Error(err)
	}
}

func TestIPsListPool(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.IPs().ListPools(); err != nil {
		t.Error(err)
	}
}

func TestIPsPoolsLifecycle(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	pool := "test-pool"
	m.IPs().DeletePool(pool)

	if _, err := m.IPs().CreatePool(pool); err != nil {
		t.Errorf("failed to create pool. %s", err)
		return
	}

	if _, err := m.IPs().PoolInfo(pool); err != nil {
		t.Errorf("failed to retrieve pool info. %s", err)
		return
	}

	if _, err := m.IPs().DeletePool(pool); err != nil {
		t.Errorf("failed to delete pool info. %s", err)
	}
}
