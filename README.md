# Mandrill [![Build Status](https://travis-ci.org/jimtsao/mandrill.svg)](https://travis-ci.org/jimtsao/mandrill)
Mandrill is a Go API for Mandrill Email Service

It supports all API calls (users, messages, tags, whitelists webhooks etc). The code is written to be as similar as possible to official libraries. Multiple API keys can be used concurrently.

### Import
    import "github.com/jimtsao/mandrill"

### Usage
Create a new Mandrill with given API Key

    m := mandrill.NewMandrill("your-api-key")

Choose an API domain

	m.Messages()
	m.Tags()
	m.Subaccounts()
	m.Whitelists()
	m.Webhooks()

Call a function associated with that domain

	msg := &Message{
		FromName:  "John Doe",
		FromEmail: TestFromEmail,
		To: []Recipient{
			{"jane@example.com", "Jane Doe", "to"},
			{"jessica@example.com", "Jessica Doe", "cc"},
		},
		Subject: "Confirmation of Membership",
		Text:    "Congratulations on becoming a member.",
	}
	response, err := m.Messages().Send(msg, false, "", nil)
	if err != nil {
		// handle error
	}

Type assert error for more details

	if ae, ok := err.(*APIError); ok {
		switch ae.Name {
		case "Invalid_Key":
		case "ValidationError":
		...
		default:
		}
	}


### Testing
Set environmental variables:

* `MANDRILL_TEST_API_KEY`
* `MANDRILL_TEST_FROM_EMAIL`

From email should have SPF and DKIM validated domain