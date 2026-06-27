package listmonk

import "time"

type Subscriber struct {
	ID        int              `json:"id,omitempty"`
	UUID      string           `json:"uuid,omitempty"`
	Email     string           `json:"email,omitempty"`
	Name      string           `json:"name,omitempty"`
	Status    string           `json:"status,omitempty"`
	Attribs   map[string]any   `json:"attribs,omitempty"`
	Lists     []SubscriberList `json:"lists,omitempty"`
	CreatedAt string           `json:"created_at,omitempty"`
	UpdatedAt string           `json:"updated_at,omitempty"`
}

type SubscriberList struct {
	ID                 int      `json:"id,omitempty"`
	UUID               string   `json:"uuid,omitempty"`
	Name               string   `json:"name,omitempty"`
	Type               string   `json:"type,omitempty"`
	Tags               []string `json:"tags,omitempty"`
	SubscriptionStatus string   `json:"subscription_status,omitempty"`
	CreatedAt          string   `json:"created_at,omitempty"`
	UpdatedAt          string   `json:"updated_at,omitempty"`
}

type SubscriberProfile struct {
	ID        int            `json:"id,omitempty"`
	UUID      string         `json:"uuid,omitempty"`
	Email     string         `json:"email,omitempty"`
	Name      string         `json:"name,omitempty"`
	Status    string         `json:"status,omitempty"`
	Attribs   map[string]any `json:"attribs,omitempty"`
	CreatedAt string         `json:"created_at,omitempty"`
	UpdatedAt string         `json:"updated_at,omitempty"`
}

type SubscriberData struct {
	Email         string              `json:"email,omitempty"`
	Profile       []SubscriberProfile `json:"profile,omitempty"`
	Subscriptions []Subscription      `json:"subscriptions,omitempty"`
	CampaignViews []map[string]any    `json:"campaign_views,omitempty"`
	LinkClicks    []map[string]any    `json:"link_clicks,omitempty"`
}

type Subscription struct {
	Name               string `json:"name,omitempty"`
	Type               string `json:"type,omitempty"`
	SubscriptionStatus string `json:"subscription_status,omitempty"`
	CreatedAt          string `json:"created_at,omitempty"`
}

type List struct {
	ID              int      `json:"id,omitempty"`
	UUID            string   `json:"uuid,omitempty"`
	Name            string   `json:"name,omitempty"`
	Type            string   `json:"type,omitempty"`
	Optin           string   `json:"optin,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	Description     string   `json:"description,omitempty"`
	SubscriberCount int      `json:"subscriber_count,omitempty"`
	CreatedAt       string   `json:"created_at,omitempty"`
	UpdatedAt       string   `json:"updated_at,omitempty"`
}

type CampaignContentType string

const (
	ContentTypeRichtext CampaignContentType = "richtext"
	ContentTypeHTML     CampaignContentType = "html"
	ContentTypeMarkdown CampaignContentType = "markdown"
	ContentTypePlain    CampaignContentType = "plain"
)

type CampaignType string

const (
	CampaignTypeRegular CampaignType = "regular"
	CampaignTypeOptin   CampaignType = "optin"
)

type CampaignStatus string

const (
	CampaignStatusDraft     CampaignStatus = "draft"
	CampaignStatusScheduled CampaignStatus = "scheduled"
	CampaignStatusRunning   CampaignStatus = "running"
	CampaignStatusPaused    CampaignStatus = "paused"
	CampaignStatusCancelled CampaignStatus = "cancelled"
	CampaignStatusFinished  CampaignStatus = "finished"
)

type CampaignList struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Campaign struct {
	ID          int                 `json:"id,omitempty"`
	UUID        string              `json:"uuid,omitempty"`
	Name        string              `json:"name,omitempty"`
	Subject     string              `json:"subject,omitempty"`
	FromEmail   string              `json:"from_email,omitempty"`
	Body        string              `json:"body,omitempty"`
	Type        CampaignType        `json:"type,omitempty"`
	ContentType CampaignContentType `json:"content_type,omitempty"`
	Status      CampaignStatus      `json:"status,omitempty"`
	Messenger   string              `json:"messenger,omitempty"`
	Tags        []string            `json:"tags,omitempty"`
	Lists       []CampaignList      `json:"lists,omitempty"`
	TemplateID  int                 `json:"template_id,omitempty"`
	SendAt      string              `json:"send_at,omitempty"`
	StartedAt   string              `json:"started_at,omitempty"`
	ToSend      int                 `json:"to_send,omitempty"`
	Sent        int                 `json:"sent,omitempty"`
	Views       int                 `json:"views,omitempty"`
	Clicks      int                 `json:"clicks,omitempty"`
	CreatedAt   string              `json:"created_at,omitempty"`
	UpdatedAt   string              `json:"updated_at,omitempty"`
}

type TemplateType string

const (
	TemplateTypeCampaign       TemplateType = "campaign"
	TemplateTypeCampaignVisual TemplateType = "campaign_visual"
	TemplateTypeTx             TemplateType = "tx"
)

type Template struct {
	ID         int          `json:"id,omitempty"`
	Name       string       `json:"name,omitempty"`
	Type       TemplateType `json:"type,omitempty"`
	Body       string       `json:"body,omitempty"`
	BodySource string       `json:"body_source,omitempty"`
	Subject    string       `json:"subject,omitempty"`
	IsDefault  bool         `json:"is_default,omitempty"`
	CreatedAt  string       `json:"created_at,omitempty"`
	UpdatedAt  string       `json:"updated_at,omitempty"`
}

type MediaFileObject struct {
	ID          int            `json:"id,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	Filename    string         `json:"filename,omitempty"`
	ContentType string         `json:"content_type,omitempty"`
	Provider    string         `json:"provider,omitempty"`
	URI         string         `json:"uri,omitempty"`
	URL         string         `json:"url,omitempty"`
	ThumbURI    string         `json:"thumb_uri,omitempty"`
	ThumbURL    string         `json:"thumb_url,omitempty"`
	Meta        map[string]any `json:"meta,omitempty"`
	CreatedAt   string         `json:"created_at,omitempty"`
}

type SMTPSettings struct {
	UUID          string           `json:"uuid,omitempty"`
	Enabled       bool             `json:"enabled,omitempty"`
	Host          string           `json:"host,omitempty"`
	Port          int              `json:"port,omitempty"`
	Username      string           `json:"username,omitempty"`
	AuthProtocol  string           `json:"auth_protocol,omitempty"`
	TLSType       string           `json:"tls_type,omitempty"`
	TLSSkipVerify bool             `json:"tls_skip_verify,omitempty"`
	MaxConns      int              `json:"max_conns,omitempty"`
	MaxMsgRetries int              `json:"max_msg_retries,omitempty"`
	IdleTimeout   string           `json:"idle_timeout,omitempty"`
	WaitTimeout   string           `json:"wait_timeout,omitempty"`
	HelloHostname string           `json:"hello_hostname,omitempty"`
	EmailHeaders  []map[string]any `json:"email_headers,omitempty"`
}

type MailBoxBounces struct {
	UUID          string `json:"uuid,omitempty"`
	Type          string `json:"type,omitempty"`
	Enabled       bool   `json:"enabled,omitempty"`
	Host          string `json:"host,omitempty"`
	Port          int    `json:"port,omitempty"`
	AuthProtocol  string `json:"auth_protocol,omitempty"`
	Username      string `json:"username,omitempty"`
	ReturnPath    string `json:"return_path,omitempty"`
	ScanInterval  string `json:"scan_interval,omitempty"`
	TLSEnabled    bool   `json:"tls_enabled,omitempty"`
	TLSSkipVerify bool   `json:"tls_skip_verify,omitempty"`
}

type Settings struct {
	SiteName        string           `json:"app.site_name,omitempty"`
	RootURL         string           `json:"app.root_url,omitempty"`
	LogoURL         string           `json:"app.logo_url,omitempty"`
	FaviconURL      string           `json:"app.favicon_url,omitempty"`
	FromEmail       string           `json:"app.from_email,omitempty"`
	NotifyEmails    []string         `json:"app.notify_emails,omitempty"`
	Lang            string           `json:"app.lang,omitempty"`
	Concurrency     int              `json:"app.concurrency,omitempty"`
	MessageRate     int              `json:"app.message_rate,omitempty"`
	BatchSize       int              `json:"app.batch_size,omitempty"`
	MaxSendErrors   int              `json:"app.max_send_errors,omitempty"`
	SMTP            []SMTPSettings   `json:"smtp,omitempty"`
	BounceMailboxes []MailBoxBounces `json:"bounce.mailboxes,omitempty"`
	BounceEnabled   bool             `json:"bounce.enabled,omitempty"`
	BounceCount     int              `json:"bounce.count,omitempty"`
}

type BounceRecord struct {
	ID             int                   `json:"id,omitempty"`
	Type           string                `json:"type,omitempty"`
	Source         string                `json:"source,omitempty"`
	Email          string                `json:"email,omitempty"`
	SubscriberID   int                   `json:"subscriber_id,omitempty"`
	SubscriberUUID string                `json:"subscriber_uuid,omitempty"`
	CampaignUUID   string                `json:"campaign_uuid,omitempty"`
	Campaign       *BounceRecordCampaign `json:"campaign,omitempty"`
	Meta           map[string]any        `json:"meta,omitempty"`
	Total          int                   `json:"total,omitempty"`
	CreatedAt      string                `json:"created_at,omitempty"`
}

type BounceRecordCampaign struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CampaignAnalyticsCount struct {
	CampaignID int        `json:"campaign_id,omitempty"`
	Count      int        `json:"count,omitempty"`
	Timestamp  *time.Time `json:"timestamp,omitempty"`
}

type CampaignStats struct {
	ID        int    `json:"id,omitempty"`
	Status    string `json:"status,omitempty"`
	ToSend    int    `json:"to_send,omitempty"`
	Sent      int    `json:"sent,omitempty"`
	Rate      int    `json:"rate,omitempty"`
	NetRate   int    `json:"net_rate,omitempty"`
	StartedAt string `json:"started_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
