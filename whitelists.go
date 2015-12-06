package mandrill

import (
	"encoding/json"
)

type Whitelists struct {
	m *Mandrill
}

// Add an email to your email rejection whitelist. If the address is currently
// on your blacklist, that blacklist entry will be removed automatically.
// comment: an optional description of why the email was whitelisted maxlength(255)
func (w *Whitelists) Add(email string, comment string) (whitelistsAddResponse, error) {
	var ret whitelistsAddResponse
	data := struct {
		APIKey  string `json:"key"`
		Email   string `json:"email"`
		Comment string `json:"comment,omitempty"`
	}{w.m.APIKey, email, comment}
	body, err := w.m.execute("/whitelists/add.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type whitelistsAddResponse struct {
	Email string `json:"email"`
	Added bool   `json:"added"`
}

func (w *Whitelists) Delete(email string) (whitelistsDeleteResponse, error) {
	var ret whitelistsDeleteResponse
	data := struct {
		APIKey string `json:"key"`
		Email  string `json:"email"`
	}{w.m.APIKey, email}
	body, err := w.m.execute("/whitelists/delete.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type whitelistsDeleteResponse struct {
	Email   string `json:"email"`
	Deleted bool   `json:"deleted"`
}

// List retrieves up to 1000 of your email rejection whitelist
// providee an email address or search prefix to limit the results
func (w *Whitelists) List(email string) ([]whitelistsListResponse, error) {
	var ret []whitelistsListResponse
	data := struct {
		APIKey string `json:"key"`
		Email  string `json:"email,omitempty"`
	}{w.m.APIKey, email}
	body, err := w.m.execute("/whitelists/list.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type whitelistsListResponse struct {
	// email that is whitelisted
	Email string `json:"email"`

	// description of why the email was whitelisted
	Detail string `json:"detail"`

	// when the email was added to the whitelist
	CreatedAt string `json:"created_at"`
}
