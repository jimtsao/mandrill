package mandrill

import (
	"testing"
)

func TestSubaccountsList(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Subaccounts().List(""); err != nil {
		t.Error(err)
	}
}

func TestSubaccountsLifecycle(t *testing.T) {
	// add subaccount
	m := NewMandrill(TestAPIKey)
	id := "test-cust-52109"
	if _, err := m.Subaccounts().Add(id, "", "", 0); err != nil {
		t.Errorf("failed to add subaccount. %s", err)
		return
	}

	// update subaccount
	if _, err := m.Subaccounts().Update(id, "", "", 0); err != nil {
		t.Errorf("failed to update subaccount. %s", err)
		return
	}

	// pause subaccount
	if _, err := m.Subaccounts().Pause(id); err != nil {
		t.Errorf("failed to pause subaccount. %s", err)
		return
	}

	// resume subaccount
	if _, err := m.Subaccounts().Resume(id); err != nil {
		t.Errorf("failed to resume subaccount. %s", err)
		return
	}

	// retrieve info
	if _, err := m.Subaccounts().Info(id); err != nil {
		t.Errorf("failed to retrieve subaccount info. %s", err)
		return
	}

	// delete subaccount
	if _, err := m.Subaccounts().Delete(id); err != nil {
		t.Errorf("failed to delete subaccount. %s", err)
		return
	}
}
