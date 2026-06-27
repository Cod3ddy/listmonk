package listmonk

type NewSubscriber struct {
	Email                   string         `json:"email"`
	Name                    string         `json:"name,omitempty"`
	Status                  string         `json:"status,omitempty"`
	Lists                   []int          `json:"lists,omitempty"`
	ListUUIDs               []string       `json:"list_uuids,omitempty"`
	Attribs                 map[string]any `json:"attribs,omitempty"`
	PreconfirmSubscriptions bool           `json:"preconfirm_subscriptions,omitempty"`
}

type UpdateSubscriber struct {
	Email                   string         `json:"email,omitempty"`
	Name                    string         `json:"name,omitempty"`
	Status                  string         `json:"status,omitempty"`
	Lists                   []int          `json:"lists,omitempty"`
	ListUUIDs               []string       `json:"list_uuids,omitempty"`
	Attribs                 map[string]any `json:"attribs,omitempty"`
	PreconfirmSubscriptions bool           `json:"preconfirm_subscriptions,omitempty"`
}

type SubscriberAction string

const (
	SubscriberActionAdd         SubscriberAction = "add"
	SubscriberActionRemove      SubscriberAction = "remove"
	SubscriberActionUnsubscribe SubscriberAction = "unsubscribe"
)

type SubscriberStatus string

const (
	SubscriberStatusConfirmed    SubscriberStatus = "confirmed"
	SubscriberStatusUnconfirmed  SubscriberStatus = "unconfirmed"
	SubscriberStatusUnsubscribed SubscriberStatus = "unsubscribed"
)

type SubscriberQueryRequest struct {
	Query         string           `json:"query,omitempty"`
	IDs           []int            `json:"ids,omitempty"`
	Action        SubscriberAction `json:"action,omitempty"`
	TargetListIDs []int            `json:"target_list_ids,omitempty"`
	Status        SubscriberStatus `json:"status,omitempty"`
}

type ListType string

const (
	ListTypePublic  ListType = "public"
	ListTypePrivate ListType = "private"
)

type ListOptin string

const (
	ListOptinSingle ListOptin = "single"
	ListOptinDouble ListOptin = "double"
)

type NewList struct {
	Name        string    `json:"name"`
	Type        ListType  `json:"type,omitempty"`
	Optin       ListOptin `json:"optin,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	Description string    `json:"description,omitempty"`
}

type CampaignRequest struct {
	Name        string              `json:"name"`
	Subject     string              `json:"subject"`
	FromEmail   string              `json:"from_email,omitempty"`
	Lists       []int               `json:"lists"`
	Type        CampaignType        `json:"type,omitempty"`
	ContentType CampaignContentType `json:"content_type,omitempty"`
	Messenger   string              `json:"messenger,omitempty"`
	TemplateID  int                 `json:"template_id,omitempty"`
	Tags        []string            `json:"tags,omitempty"`
	SendLater   bool                `json:"send_later,omitempty"`
	SendAt      string              `json:"send_at,omitempty"`
	Headers     []map[string]string `json:"headers,omitempty"`
}

type CampaignUpdate struct {
	Name              string              `json:"name,omitempty"`
	Subject           string              `json:"subject,omitempty"`
	FromEmail         string              `json:"from_email,omitempty"`
	Lists             []int               `json:"lists,omitempty"`
	Type              CampaignType        `json:"type,omitempty"`
	ContentType       CampaignContentType `json:"content_type,omitempty"`
	Body              string              `json:"body,omitempty"`
	AltBody           string              `json:"altbody,omitempty"`
	Messenger         string              `json:"messenger,omitempty"`
	TemplateID        int                 `json:"template_id,omitempty"`
	Tags              []string            `json:"tags,omitempty"`
	SendLater         bool                `json:"send_later,omitempty"`
	SendAt            map[string]any      `json:"send_at,omitempty"`
	Headers           []map[string]string `json:"headers,omitempty"`
	Archive           bool                `json:"archive,omitempty"`
	ArchiveTemplateID int                 `json:"archive_template_id,omitempty"`
	ArchiveMeta       map[string]any      `json:"archive_meta,omitempty"`
}

type CampaignArchiveRequest struct {
	Archive           bool           `json:"archive,omitempty"`
	ArchiveTemplateID int            `json:"archive_template_id,omitempty"`
	ArchiveMeta       map[string]any `json:"archive_meta,omitempty"`
}

type CampaignStatusRequest struct {
	Status CampaignStatus `json:"status"`
}

type CampaignContentRequest struct {
	Body        string              `json:"body,omitempty"`
	ContentType CampaignContentType `json:"content_type,omitempty"`
	TemplateID  int                 `json:"template_id,omitempty"`
}

type NewTemplate struct {
	Name       string       `json:"name"`
	Type       TemplateType `json:"type"`
	Body       string       `json:"body"`
	BodySource string       `json:"body_source,omitempty"`
	Subject    string       `json:"subject,omitempty"`
}

type UpdateTemplate struct {
	Name       string       `json:"name,omitempty"`
	Type       TemplateType `json:"type,omitempty"`
	Body       string       `json:"body,omitempty"`
	BodySource string       `json:"body_source,omitempty"`
	Subject    string       `json:"subject,omitempty"`
}

type TransactionalMessage struct {
	SubscriberID    int                 `json:"subscriber_id,omitempty"`
	SubscriberEmail string              `json:"subscriber_email,omitempty"`
	TemplateID      int                 `json:"template_id"`
	FromEmail       string              `json:"from_email,omitempty"`
	Messenger       string              `json:"messenger,omitempty"`
	ContentType     CampaignContentType `json:"content_type,omitempty"`
	Headers         []map[string]any    `json:"headers,omitempty"`
	Data            map[string]any      `json:"data,omitempty"`
}

type SMTPTest struct {
	UUID          string           `json:"uuid,omitempty"`
	Enabled       bool             `json:"enabled,omitempty"`
	Host          string           `json:"host,omitempty"`
	Port          int              `json:"port,omitempty"`
	Username      string           `json:"username,omitempty"`
	Password      string           `json:"password,omitempty"`
	AuthProtocol  string           `json:"auth_protocol,omitempty"`
	TLSType       string           `json:"tls_type,omitempty"`
	TLSSkipVerify bool             `json:"tls_skip_verify,omitempty"`
	MaxConns      int              `json:"max_conns,omitempty"`
	MaxMsgRetries int              `json:"max_msg_retries,omitempty"`
	IdleTimeout   string           `json:"idle_timeout,omitempty"`
	WaitTimeout   string           `json:"wait_timeout,omitempty"`
	HelloHostname string           `json:"hello_hostname,omitempty"`
	EmailHeaders  []map[string]any `json:"email_headers,omitempty"`
	Email         string           `json:"email,omitempty"`
}

type SortOrder string

const (
	SortAsc  SortOrder = "ASC"
	SortDesc SortOrder = "DESC"
)

type GetSubscribersParams struct {
	Page               int       `form:"page"`
	PerPage            int       `form:"per_page"`
	Query              string    `form:"query"`
	OrderBy            string    `form:"order_by"`
	Order              SortOrder `form:"order"`
	ListID             []int     `form:"list_id"`
	SubscriptionStatus string    `form:"subscription_status"`
}

type GetListsParams struct {
	Page    int       `form:"page"`
	PerPage int       `form:"per_page"`
	Query   string    `form:"query"`
	OrderBy string    `form:"order_by"`
	Order   SortOrder `form:"order"`
	Minimal bool      `form:"minimal"`
	Tag     []string  `form:"tag"`
}

type GetCampaignsParams struct {
	Page    int              `form:"page"`
	PerPage int              `form:"per_page"`
	Query   string           `form:"query"`
	Status  []CampaignStatus `form:"status"`
	Tags    []string         `form:"tags"`
	OrderBy string           `form:"order_by"`
	Order   SortOrder        `form:"order"`
	NoBody  bool             `form:"no_body"`
}

type GetBouncesParams struct {
	Page       int    `form:"page"`
	PerPage    int    `form:"per_page"`
	CampaignID int    `form:"campaign_id"`
	Source     string `form:"source"`
	OrderBy    string `form:"order_by"`
	Order      string `form:"order"`
}
