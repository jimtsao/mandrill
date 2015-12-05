package mandrill

import (
	"encoding/json"
	"fmt"
	"time"
)

type Messages struct {
	*Mandrill
}

func (m *Messages) Send(message *Message, async bool, ipPool string, sendAt *time.Time) ([]SendResponse, error) {
	var tsend string
	if sendAt != nil {
		tsend = ToMandrillTime(*sendAt)
	}
	data := struct {
		APIkey  string   `json:"key"`
		Message *Message `json:"message"`

		// enable a background sending mode that is optimized for bulk sending.
		// In async mode, messages/send will immediately return a status of "queued" for every recipient.
		// To handle rejections when sending in async mode, set up a webhook for the 'reject' event.
		// Defaults to false for messages with no more than 10 recipients; messages with more than 10
		// recipients are always sent asynchronously, regardless of the value of async.
		Async bool `json:"async,omitempty"`

		// the name of the dedicated ip pool that should be used to send the message.
		// If you do not have any dedicated IPs, this parameter has no effect.
		// If you specify a pool that does not exist, your default pool will be used instead.
		IPPool string `json:"ip_pool,omitempty"`
		SendAt string `json:"send_at,omitempty"`
	}{m.APIKey, message, async, ipPool, tsend}
	resp, err := m.execute("/messages/send.json", data)
	if err != nil {
		return nil, err
	}

	var ret []SendResponse
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (m *Mandrill) SimpleSend(from string, to string, subject string, body string) (SendResponse, error) {
	msg := &Message{
		FromEmail: from,
		To:        []Recipient{{Email: to}},
		Subject:   subject,
		Text:      body,
	}
	data := struct {
		APIKey  string   `json:"key"`
		Message *Message `json:"message"`
	}{m.APIKey, msg}
	resp, err := m.execute("/messages/send.json", data)
	if err != nil {
		return SendResponse{}, err
	}

	var ret []SendResponse
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return SendResponse{}, err
	} else if len(ret) != 1 {
		return SendResponse{}, fmt.Errorf("received more than one response: %+v", ret)
	}

	return ret[0], nil
}

type SendResponse struct {
	// the email address of the recipient
	Email string `json:"email"`

	// the sending status of the recipient - either "sent", "queued", "scheduled", "rejected", or "invalid"
	Status string `json:"status"`

	// the reason for the rejection if the recipient status is "rejected"
	// one of "hard-bounce", "soft-bounce", "spam", "unsub", "custom", "invalid-sender",
	// "invalid", "test-mode-limit", "unsigned", or "rule"
	RejectReason string `json:"reject_reason"`

	// the message's unique id
	Id string `json:"_id"`
}

// Message contains information on the message to send
type Message struct {
	// the full HTML content to be sent
	HTML string `json:"html,omitempty"`

	// optional full text content to be sent
	Text string `json:"text,omitempty"`

	// the message subject
	Subject string `json:"subject,omitempty"`

	// the sender email address
	FromEmail string `json:"from_email"`

	// optional from name to be used
	FromName string `json:"from_name,omitempty"`

	// an array of recipient information
	To []Recipient `json:"to"`

	// optional extra headers to add to the message (most headers are allowed)
	Headers map[string]string `json:"headers,omitempty"`

	// whether or not this message is important, and should be delivered ahead of non-important messages
	Important bool `json:"important,omitempty"`

	// whether or not to turn on open tracking for the message
	TrackOpens bool `json:"track_opens,omitempty"`

	// whether or not to turn on click tracking for the message
	TrackClicks bool `json:"track_clicks,omitempty"`

	// whether or not to automatically generate a text part for messages that are not given text
	AutoText bool `json:"auto_text,omitempty"`

	// whether or not to automatically generate an HTML part for messages that are not given HTML
	AutoHTML bool `json:"auto_html,omitempty"`

	// whether or not to automatically inline all CSS styles provided in the message HTML - only for
	// HTML documents less than 256KB in size
	InlineCSS bool `json:"inline_css,omitempty"`

	// whether or not to strip the query string from URLs when aggregating tracked URL data
	URLStripQueries bool `json:"url_strip_qs,omitempty"`

	// whether or not to expose all recipients in to "To" header for each email
	PreserveRecipients bool `json:"preserve_recipients,omitempty"`

	// set to false to remove content logging for sensitive emails
	ViewContentLink bool `json:"view_content_link,omitempty"`

	// an optional address to receive an exact copy of each recipient's email
	BCCAddress string `json:"bcc_address,omitempty"`

	// a custom domain to use for tracking opens and clicks instead of mandrillapp.com
	TrackingDomain string `json:"tracking_domain,omitempty"`

	// a custom domain to use for SPF/DKIM signing instead of mandrill
	// (for "via" or "on behalf of" in email clients)
	SigningDomain string `json:"signing_domain,omitempty"`

	// a custom domain to use for the messages's return-path
	ReturnPathDomain string `json:"return_path_domain,omitempty"`

	// whether to evaluate merge tags in the message. Will automatically be set to true if either
	// merge_vars or global_merge_vars are provided.
	Merge bool `json:"merge,omitempty"`

	// the merge tag language to use when evaluating merge tags, either mailchimp or handlebars
	// oneof(mailchimp, handlebars)
	MergeLang string `json:"merge_language,omitempty"`

	// global merge variables to use for all recipients. You can override these per recipient
	GlobalMergeVars []MergeVar `json:"global_merge_vars,omitempty"`

	// per-recipient merge variables, which override global merge variables with the same name
	MergeVars []RecipientMergeVar `json:"merge_vars,omitempty"`

	// an array of string to tag the message with. Stats are accumulated using tags,
	// though we only store the first 100 we see, so this should not be unique or change frequently.
	// Tags should be 50 characters or less. Any tags starting with an underscore are reserved for
	// internal use and will cause errors.
	Tags []string `json:"tags,omitempty"`

	// the unique id of a subaccount for this message - must already exist or will fail with an error
	SubAccount string `json:"subaccount,omitempty"`

	// an array of strings indicating for which any matching URLs will automatically have Google Analytics
	// parameters appended to their query string automatically
	GoogleAnalyticsDomain []string `json:"google_analytics_domains,omitempty"`

	// optional string indicating the value to set for the utm_campaign tracking parameter.
	// If this isn't provided the email's from address will be used instead
	GoogleAnalyticsCampaign []string `json:"google_analytics_campaign,omitempty"`

	// metadata an associative array of user metadata. Mandrill will store this metadata and make it
	// available for retrieval. In addition, you can select up to 10 metadata fields to index and make
	// searchable using the Mandrill search api
	Metadata map[string]string `json:"metadata,omitempty"`

	// Per-recipient metadata that will override the global values specified in the metadata parameter
	RecipientMetadata []RecipientMetadata `json:"recipient_metadata,omitempty"`

	// an array of supported attachments to add to the message
	Attachments []*Attachment `json:"attachments,omitempty"`

	// an array of embedded images to add to the message
	Images []*Image `json:"images,omitempty"`
}

type Recipient struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
	// "to" (default), "cc" or "bcc"
	Type string `json:"type,omitempty"`
}

type MergeVar struct {
	// Merge tags can be composed of alphanumeric characters and underscores
	// Case insensitive. May not start with _ or contain :
	Name string `json:"name,omitempty"`

	// the merge variable's content
	Content interface{} `json:"content,omitempty"`
}

type RecipientMergeVar struct {
	// the email address of the recipient that the merge variables should apply to
	Recipient string     `json:"rcpt"`
	Vars      []MergeVar `json:"vars,omitempty"`
}

type RecipientMetadata struct {
	Recipient string
	Values    map[string]string
}

type Attachment struct {
	// the file name of the attachment
	Name string `json:"name,omitempty"`

	// the MIME type of the attachment
	Type string `json:"type,omitempty"`

	// the content of the attachment as a base64-encoded string
	Content string `json:"content,omitempty"`
}

type Image struct {
	// the Content ID of the image - use <img src="cid:THIS_VALUE"> to reference the image in your HTML content
	Name string `json:"name,omitempty"`

	// the MIME type of the image - must start with "image/"
	Type string `json:"type,omitempty"`

	// the content of the image as a base64-encoded string
	Content string `json:"content,omitempty"`
}
