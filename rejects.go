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
func (r *Rejects) Add(email string, comment string, subaccount string) (rejectsAddResponse, error) {
	var ret rejectsAddResponse
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

// rejectsAddResponse contains the address and the result of the operation
type rejectsAddResponse struct {
	// the email address you provided
	Email string `json:"email"`

	// whether the operation succeeded
	Added bool `json:"added"`
}

// Delete an email rejection. There is no limit to how many rejections you can remove from your blacklist
// however each deletion has an effect on your reputation
func (r *Rejects) Delete(email string, subaccount string) (rejectsDeleteResponse, error) {
	var ret rejectsDeleteResponse
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

// rejectsDeleteResponse contains the address and whether the deletion succeeded
type rejectsDeleteResponse struct {
	// the email address that was removed from the blacklist
	Email string `json:"email"`

	// whether the address was deleted successfully.
	Deleted bool `json:"deleted"`

	// the subaccount blacklist that the address was removed from, if any
	Subaccount string `json:"subaccount"`
}

// List retrieves up to 1000 rejection entries
func (r *Rejects) List(email string, expired bool, subaccount string) ([]rejectsListResponse, error) {
	var ret []rejectsListResponse
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

// rejectsListResponse contains information for each rejection blacklist entry
type rejectsListResponse struct {
	// the email that is blocked
	Email string `json:"email"`

	// the type of event (hard-bounce, soft-bounce, spam, unsub, custom) that caused this rejection
	Reason string `json:"reason"`

	// extended details about the event, such as the SMTP diagnostic
	// for bounces or the comment for manually-created rejections
	Detail string `json:"detail"`

	// when the email was added to the blacklist
	CreatedAt string `json:"created_at"`

	// the timestamp of the most recent event that either created or renewed this rejection
	LastEventAt string `json:"last_event_at"`

	// when the blacklist entry will expire (this may be in the past)
	ExpiresAt string `json:"expires_at"`

	// whether the blacklist entry has expired
	Expired bool `json:"expired"`

	// the subaccount that this blacklist entry applies to, or null if none
	Subaccount string `json:"subaccount"`

	// the sender that this blacklist entry applies to, or null if none
	Sender *struct {
		// the sender's email address
		Address string `json:"address"`

		// the date and time that the sender was first seen by Mandrill as a
		// UTC date string in YYYY-MM-DD HH:MM:SS format
		CreatedAt string `json:"created_at"`

		// the total number of messages sent by this sender
		Sent int `json:"sent"`

		// the total number of hard bounces by messages by this sender
		HardBounces int `json:"hard_bounces"`

		// the total number of soft bounces by messages by this sender
		SoftBounces int `json:"soft_bounces"`

		// the total number of rejected messages by this sender
		Rejects int `json:"rejects"`

		// the total number of spam complaints received for messages by this sender
		Complaints int `json:"complaints"`

		// the total number of unsubscribe requests received for messages by this sender
		Unsubs int `json:"unsubs"`

		// the total number of times messages by this sender have been opened
		Opens int `json:"opens"`

		// the total number of times tracked URLs in messages by this sender have been clicked
		Clicks int `json:"clicks"`

		// the number of unique opens for emails sent for this sender
		UniqueOpens int `json:"unique_opens"`

		// the number of unique clicks for emails sent for this sender
		UniqueClicks int `json:"unique_clicks"`
	} `json:"sender, omitempty"`
}
