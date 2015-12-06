package mandrill

import (
	"encoding/json"
)

type Rejects struct {
	m *Mandrill
}

// Add an email to your email rejection blacklist
// email: an email address to block
// comment: an optional comment describing the rejection
// subaccount:  an optional unique identifier for the subaccount to limit the blacklist entry maxlength(255)
func (r *Rejects) Add(email string, comment string, subaccount string) (rejectAddResponse, error) {
	var ret rejectAddResponse
	data := struct {
		APIKey     string `json:"key"`
		Email      string `json:"email"`
		Comment    string `json:"comment,omitempty"`
		Subaccount string `json:"subaccount,omitempty"`
	}{r.m.APIKey, email, comment, subaccount}
	body, err := r.m.execute("/rejects/add.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// rejectAddResponse contains the address and the result of the operation
type rejectAddResponse struct {
	// the email address you provided
	Email string `json:"email"`

	// whether the operation succeeded
	Added bool `json:"added"`
}

// Delete an email rejection. There is no limit to how many rejections you can remove from your blacklist
// however each deletion has an effect on your reputation
func (r *Rejects) Delete(email string, subaccount string) (rejectDeleteResponse, error) {
	var ret rejectDeleteResponse
	data := struct {
		APIKey     string `json:"key"`
		Email      string `json:"email"`
		Subaccount string `json:"subaccount,omitempty"`
	}{r.m.APIKey, email, subaccount}
	body, err := r.m.execute("/rejects/delete.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// rejectDeleteResponse contains the address and whether the deletion succeeded
type rejectDeleteResponse struct {
	// the email address that was removed from the blacklist
	Email string `json:"email"`

	// whether the address was deleted successfully.
	Deleted bool `json:"deleted"`

	// the subaccount blacklist that the address was removed from, if any
	Subaccount string `json:"subaccount"`
}

// List retrieves up to 1000 rejection entries
func (r *Rejects) List(email string, expired bool, subaccount string) ([]rejectListResponse, error) {
	var ret []rejectListResponse
	data := struct {
		APIKey         string `json:"key"`
		Email          string `json:"email,omitempty"`
		IncludeExpired bool   `json:"include_expired,omitempty"`
		Subaccount     string `json:"subaccount,omitempty`
	}{r.m.APIKey, email, expired, subaccount}
	body, err := r.m.execute("/rejects/list.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// rejectListResponse contains information for each rejection blacklist entry
type rejectListResponse struct {
	Email       string `json:"email"`
	Reason      string `json:"reason"`
	Detail      string `json:"detail"`
	CreatedAt   string `json:"created_at"`
	LastEventAt string `json:"last_event_at"`
	ExpiresAt   string `json:"expires_at"`
	Expired     bool   `json:"expired"`
	Subaccount  string `json:"subaccount"`
	Sender      struct {
		Address      string `json:"address"`
		CreatedAt    string `json:"created_at"`
		Sent         int    `json:"sent"`
		HardBounces  int    `json:"hard_bounces"`
		SoftBounces  int    `json:"soft_bounces"`
		Rejects      int    `json:"rejects"`
		Complaints   int    `json:"complaints"`
		Unsubs       int    `json:"unsubs"`
		Opens        int    `json:"opens"`
		Clicks       int    `json:"clicks"`
		UniqueOpens  int    `json:"unique_opens"`
		UniqueClicks int    `json:"unique_clicks"`
	} `json:"sender"`
}
