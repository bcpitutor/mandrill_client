package main

import (
	"net/http"
)

// Client manages requests to the Mandrill API
type Client struct {
	// mandrill API key
	Key string
	// Mandrill API base. e.g. "https://mandrillapp.com/api/1.0/"
	BaseURL string
	// Requests are transported through this client
	HTTPClient *http.Client
}

type Message struct {
	// the full HTML content to be sent
	HTML string `json:"html,omitempty"`
	// optional full text content to be sent
	Text string `json:"text,omitempty"`
	// the message subject
	Subject string `json:"subject,omitempty"`
	// the sender email address.
	FromEmail string `json:"from_email,omitempty"`
	// optional from name to be used
	FromName string `json:"from_name,omitempty"`
	// an array of recipient information.
	To []*To `json:"to"`
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
	// whether or not to automatically inline all CSS styles provided in the message HTML - only for HTML documents less than 256KB in size
	InlineCSS bool `json:"inline_css,omitempty"`
	// whether or not to strip the query string from URLs when aggregating tracked URL data
	URLStripQS bool `json:"url_strip_qs,omitempty"`
	// whether or not to expose all recipients in to "To" header for each email
	PreserveRecipients bool `json:"preserve_recipients,omitempty"`
	// set to false to remove content logging for sensitive emails
	ViewContentLink bool `json:"view_content_link,omitempty"`
	// an optional address to receive an exact copy of each recipient's email
	BCCAddress string `json:"bcc_address,omitempty"`
	// a custom domain to use for tracking opens and clicks instead of mandrillapp.com
	TrackingDomain string `json:"tracking_domain,omitempty"`
	// a custom domain to use for SPF/DKIM signing instead of mandrill (for "via" or "on behalf of" in email clients)
	SigningDomain string `json:"signing_domain,omitempty"`
	// a custom domain to use for the messages's return-path
	ReturnPathDomain string `json:"return_path_domain,omitempty"`
	// whether to evaluate merge tags in the message. Will automatically be set to true if either merge_vars or global_merge_vars are provided.
	Merge bool `json:"merge,omitempty"`
	// the merge tag language to use when evaluating merge tags, either mailchimp or handlebars
	MergeLanguage string `json:"merge_language,omitempty"`
	// global merge variables to use for all recipients. You can override these per recipient.
	GlobalMergeVars []*Variable `json:"global_merge_vars,omitempty"`
	// per-recipient merge variables, which override global merge variables with the same name.
	MergeVars []*RcptMergeVars `json:"merge_vars,omitempty"`
	// an array of string to tag the message with. Stats are accumulated using tags, though we only store the first 100 we see, so this should not be unique or change frequently. Tags should be 50 characters or less. Any tags starting with an underscore are reserved for internal use and will cause errors.
	Tags []string `json:"tags,omitempty"`
	// the unique id of a subaccount for this message - must already exist or will fail with an error
	Subaccount string `json:"subaccount,omitempty"`
	// an array of strings indicating for which any matching URLs will automatically have Google Analytics parameters appended to their query string automatically.
	GoogleAnalyticsDomains []string `json:"google_analytics_domains,omitempty"`
	// optional string indicating the value to set for the utm_campaign tracking parameter. If this isn't provided the email's from address will be used instead.
	GoogleAnalyticsCampaign string `json:"google_analytics_campaign,omitempty"`
	// metadata an associative array of user metadata. Mandrill will store this metadata and make it available for retrieval. In addition, you can select up to 10 metadata fields to index and make searchable using the Mandrill search api.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Per-recipient metadata that will override the global values specified in the metadata parameter.
	RecipientMetadata []*RcptMetadata `json:"recipient_metadata,omitempty"`
	// an array of supported attachments to add to the message
	Attachments []*Attachment `json:"attachments,omitempty"`
	// an array of embedded images to add to the message
	Images []*Attachment `json:"images,omitempty"`
	// enable a background sending mode that is optimized for bulk sending. In async mode, messages/send will immediately return a status of "queued" for every recipient. To handle rejections when sending in async mode, set up a webhook for the 'reject' event. Defaults to false for messages with no more than 10 recipients; messages with more than 10 recipients are always sent asynchronously, regardless of the value of async.
	Async bool `json:"-"`
	// the name of the dedicated ip pool that should be used to send the message. If you do not have any dedicated IPs, this parameter has no effect. If you specify a pool that does not exist, your default pool will be used instead.
	IPPool string `json:"-"`
	// when this message should be sent as a UTC timestamp in YYYY-MM-DD HH:MM:SS format. If you specify a time in the past, the message will be sent immediately. An additional fee applies for scheduled email, and this feature is only available to accounts with a positive balance.
	SendAt string `json:"-"`
}

// To is a single recipient's information.
type To struct {
	// the email address of the recipient
	Email string `json:"email"`
	// the optional display name to use for the recipient
	Name string `json:"name,omitempty"`
	// the header type to use for the recipient, defaults to "to" if not provided
	// oneof(to, cc, bcc)
	Type string `json:"type,omitempty"`
}

// Variable is key/value data used throughout the Mandrill API
type Variable struct {
	Name    string      `json:"name"`
	Content interface{} `json:"content"`
}

// RcptMergeVars holds per-recipient merge variables
type RcptMergeVars struct {
	// the email address of the recipient that the merge variables should apply to
	Rcpt string `json:"rcpt"`
	// the recipient's merge variables
	Vars []*Variable `json:"vars"`
}

// RcptMetadata holds metadata for a single recipient
type RcptMetadata struct {
	// the email address of the recipient that the metadata is associated with
	Rcpt string `json:"rcpt"`
	// an associated array containing the recipient's unique metadata. If a key exists in both the per-recipient metadata and the global metadata, the per-recipient metadata will be used.
	Values map[string]interface{} `json:"values"`
}

// Attachment represents a single supported attachment
type Attachment struct {
	// the MIME type of the attachment
	Type string `json:"type"`
	// the file name of the attachment
	Name string `json:"name"`
	// the content of the attachment as a base64-encoded string
	Content string `json:"content"`
}

// Response holds details of the message status
type Response struct {
	// the email address of the recipient
	Email string `json:"email"`
	// the sending status of the recipient - either "sent", "queued", "scheduled", "rejected", or "invalid"
	Status string `json:"status"`
	// the reason for the rejection if the recipient status is "rejected" - one of "hard-bounce", "soft-bounce", "spam", "unsub", "custom", "invalid-sender", "invalid", "test-mode-limit", or "rule"
	RejectionReason string `json:"reject_reason"`
	// the message's unique id
	Id string `json:"_id"`
}

// Error reprents an error from the Mandrill API
// - Invalid_Key -The provided API key is not a valid Mandrill API key\r
// - PaymentRequired -The requested feature requires payment.\r
// - Unknown_Subaccount - The provided subaccount id does not exist.\r
// - ValidationError - The parameters passed to the API call are invalid or not provided when required\r
// - GeneralError - An unexpected error occurred processing the request. Mandrill developers will be   notified.\r
type Error struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}
