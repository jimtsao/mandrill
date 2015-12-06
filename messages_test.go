package mandrill

import (
	"testing"
)

func TestSendMessage(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	msg := &Message{
		FromEmail: TestFromEmail,
		To: []Recipient{
			{Email: "accept@test.mandrillapp.com"},
		},
	}
	rr, err := m.Messages().Send(msg, false, "", nil)
	if err != nil {
		t.Error(err)
		return
	}

	for _, r := range rr {
		if r.Status != "sent" && r.RejectReason != "" {
			t.Errorf("Response: %+v, Error: %s", r, err)
			return
		}
	}
}

func TestSendMessageReject(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	msg := &Message{
		FromEmail: TestFromEmail,
		To: []Recipient{
			{Email: "reject@test.mandrillapp.com"},
		},
	}
	rr, err := m.Messages().Send(msg, false, "", nil)
	if err != nil {
		t.Error(err)
		return
	}

	if len(rr) != 1 {
		t.Errorf("Expected 1 response. Received: %d. Content: %+v", len(rr), rr)
		return
	}

	if rr[0].Status != "rejected" {
		t.Errorf("Expected rejected response. Response: %+v", rr[0])
	}
}

func TestSimpleSend(t *testing.T) {
	m := NewMandrill(TestAPIKey)
	if r, err := m.Messages().SimpleSend(TestFromEmail,
		"devnull@test.mandrillapp.com",
		"Simple Send",
		"Simple Send Body"); err != nil {
		t.Errorf("Expected sent response. Response: %+v. Error: %s", r, err)
	} else if r.Status == "rejected" {
		t.Errorf("failed to send. Response: %+v", r)
	}
}
