package mandrill

import (
	"testing"
)

func TestWebhooksList(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if _, err := m.Webhooks().List(); err != nil {
		t.Error(err)
	}
}

func TestWebhooksLifecycle(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	var id int
	if resp, err := m.Webhooks().Add(
		"http://google.com",
		"My Example Webhook",
		[]string{"send", "open", "click"},
	); err != nil {
		t.Errorf("failed to add webhook. error: %s", err)
		return
	} else {
		id = resp.Id
	}

	if _, err := m.Webhooks().Update(id,
		"http://google.com",
		"Updated Webhook",
		[]string{"send", "open"},
	); err != nil {
		t.Errorf("failed to update webhook. error: %s", err)
		return
	}

	if resp, err := m.Webhooks().Info(id); err != nil {
		t.Errorf("failed to retrieve webhook info. error: %s", err)
		return
	} else if resp.Description != "Updated Webhook" {
		t.Errorf("\nexpected description \"Updated Webhook\"\nreceived: \"%s\"", resp.Description)
		return
	}

	if _, err := m.Webhooks().Delete(id); err != nil {
		t.Errorf("failed to delete webhook. error: %s", err)
		return
	}
}
