package listmonk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)


type Client struct {
	baseURL    string
	httpClient *http.Client
	username   string
	token      string
}

func New(baseURL, username, token string) *Client {
	return &Client{
		baseURL:    strings.TrimRight(baseURL, "/"),
		username:   username,
		token:      token,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) newRequest(method, path string, body any) (*http.Request, error) {
	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		r = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, c.baseURL+"/api"+path, r)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s:%s", c.username, c.token))

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

func decode[T any](c *Client, req *http.Request) (T, error) {
	var zero T
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return zero, err
	}
	defer resp.Body.Close()

	var env struct {
		Data    T      `json:"data"`
		Message string `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&env); err != nil {
		if resp.StatusCode >= 400 {
			return zero, &APIError{Code: resp.StatusCode, Message: resp.Status}
		}
		return zero, err
	}
	if resp.StatusCode >= 400 {
		return zero, &APIError{Code: resp.StatusCode, Message: env.Message}
	}
	return env.Data, nil
}

func setPage(q url.Values, page, perPage int) {
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	if perPage > 0 {
		q.Set("per_page", strconv.Itoa(perPage))
	}
}

func (c *Client) Health() error {
	req, err := c.newRequest(http.MethodGet, "/health", nil)
	if err != nil {
		return err
	}
	_, err = decode[bool](c, req)
	return err
}

func (c *Client) GetServerConfig() (*ServerConfig, error) {
	req, err := c.newRequest(http.MethodGet, "/config", nil)
	if err != nil {
		return nil, err
	}
	result, err := decode[ServerConfig](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetDashboardCounts() (*DashboardCounts, error) {
	req, err := c.newRequest(http.MethodGet, "/dashboard/counts", nil)
	if err != nil {
		return nil, err
	}
	result, err := decode[DashboardCounts](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetDashboardCharts() (*DashboardCharts, error) {
	req, err := c.newRequest(http.MethodGet, "/dashboard/charts", nil)
	if err != nil {
		return nil, err
	}
	result, err := decode[DashboardCharts](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetSubscribers(params GetSubscribersParams) (*Page[Subscriber], error) {
	req, err := c.newRequest(http.MethodGet, "/subscribers", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	setPage(q, params.Page, params.PerPage)
	if params.Query != "" {
		q.Set("query", params.Query)
	}
	if params.OrderBy != "" {
		q.Set("order_by", params.OrderBy)
	}
	if params.Order != "" {
		q.Set("order", string(params.Order))
	}
	if params.SubscriptionStatus != "" {
		q.Set("subscription_status", params.SubscriptionStatus)
	}
	for _, id := range params.ListID {
		q.Add("list_id", strconv.Itoa(id))
	}
	req.URL.RawQuery = q.Encode()
	result, err := decode[Page[Subscriber]](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetSubscriber(id int) (*Subscriber, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/subscribers/%d", id), nil)
	if err != nil {
		return nil, err
	}
	result, err := decode[Subscriber](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) CreateSubscriber(sub NewSubscriber) (*Subscriber, error) {
	req, err := c.newRequest(http.MethodPost, "/subscribers", sub)
	if err != nil {
		return nil, err
	}
	result, err := decode[Subscriber](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) UpdateSubscriber(id int, sub UpdateSubscriber) (*Subscriber, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("/subscribers/%d", id), sub)
	if err != nil {
		return nil, err
	}
	result, err := decode[Subscriber](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) DeleteSubscriber(id int) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/subscribers/%d", id), nil)
	if err != nil {
		return err
	}
	_, err = decode[bool](c, req)
	return err
}


func (c *Client) ManageSubscriberLists(body SubscriberQueryRequest) error {
	req, err := c.newRequest(http.MethodPut, "/subscribers/lists", body)
	if err != nil {
		return err
	}
	_, err = decode[bool](c, req)
	return err
}


func (c *Client) ManageSubscriberListsById(id int, body SubscriberQueryRequest) error {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("/subscribers/%d/lists", id), body)
	if err != nil {
		return err
	}
	_, err = decode[bool](c, req)
	return err
}


func (c *Client) SendSubscriberOptin(id int) error {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("/subscribers/%d/optin", id), nil)
	if err != nil {
		return err
	}
	_, err = decode[bool](c, req)
	return err
}


func (c *Client) ExportSubscriberData(id int) (*SubscriberData, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/subscribers/%d/export", id), nil)
	if err != nil {
		return nil, err
	}
	result, err := decode[SubscriberData](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) GetLists(params GetListsParams) (*Page[List], error) {
	req, err := c.newRequest(http.MethodGet, "/lists", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	setPage(q, params.Page, params.PerPage)
	if params.Query != "" {
		q.Set("query", params.Query)
	}
	if params.OrderBy != "" {
		q.Set("order_by", params.OrderBy)
	}
	if params.Order != "" {
		q.Set("order", string(params.Order))
	}
	if params.Minimal {
		q.Set("minimal", "true")
	}
	for _, tag := range params.Tag {
		q.Add("tag", tag)
	}
	req.URL.RawQuery = q.Encode()
	result, err := decode[Page[List]](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) GetList(id int) (*List, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/lists/%d", id), nil)
	if err != nil {
		return nil, err
	}
	result, err := decode[List](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) CreateList(list NewList) (*List, error) {
	req, err := c.newRequest(http.MethodPost, "/lists", list)
	if err != nil {
		return nil, err
	}
	result, err := decode[List](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) UpdateList(id int, list NewList) (*List, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("/lists/%d", id), list)
	if err != nil {
		return nil, err
	}
	result, err := decode[List](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) DeleteList(id int) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/lists/%d", id), nil)
	if err != nil {
		return err
	}
	_, err = decode[bool](c, req)
	return err
}


func (c *Client) GetCampaigns(params GetCampaignsParams) (*Page[Campaign], error) {
	req, err := c.newRequest(http.MethodGet, "/campaigns", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	setPage(q, params.Page, params.PerPage)
	if params.Query != "" {
		q.Set("query", params.Query)
	}
	for _, s := range params.Status {
		q.Add("status", string(s))
	}
	for _, t := range params.Tags {
		q.Add("tags", t)
	}
	if params.OrderBy != "" {
		q.Set("order_by", params.OrderBy)
	}
	if params.Order != "" {
		q.Set("order", string(params.Order))
	}
	if params.NoBody {
		q.Set("no_body", "true")
	}
	req.URL.RawQuery = q.Encode()
	result, err := decode[Page[Campaign]](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) GetCampaign(id int) (*Campaign, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/campaigns/%d", id), nil)
	if err != nil {
		return nil, err
	}
	result, err := decode[Campaign](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) CreateCampaign(campaign CampaignRequest) (*Campaign, error) {
	req, err := c.newRequest(http.MethodPost, "/campaigns", campaign)
	if err != nil {
		return nil, err
	}
	result, err := decode[Campaign](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) UpdateCampaign(id int, campaign CampaignUpdate) (*Campaign, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("/campaigns/%d", id), campaign)
	if err != nil {
		return nil, err
	}
	result, err := decode[Campaign](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) UpdateCampaignStatus(id int, status CampaignStatus) (*Campaign, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("/campaigns/%d/status", id), CampaignStatusRequest{Status: status})
	if err != nil {
		return nil, err
	}
	result, err := decode[Campaign](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) DeleteCampaign(id int) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/campaigns/%d", id), nil)
	if err != nil {
		return err
	}
	_, err = decode[bool](c, req)
	return err
}


func (c *Client) GetRunningCampaignStats(campaignIDs ...int) ([]CampaignStats, error) {
	req, err := c.newRequest(http.MethodGet, "/campaigns/running/stats", nil)
	if err != nil {
		return nil, err
	}
	if len(campaignIDs) > 0 {
		q := req.URL.Query()
		for _, id := range campaignIDs {
			q.Add("campaign_id", strconv.Itoa(id))
		}
		req.URL.RawQuery = q.Encode()
	}
	result, err := decode[[]CampaignStats](c, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func (c *Client) GetTemplates(noBody bool) ([]Template, error) {
	req, err := c.newRequest(http.MethodGet, "/templates", nil)
	if err != nil {
		return nil, err
	}
	if noBody {
		q := req.URL.Query()
		q.Set("no_body", "true")
		req.URL.RawQuery = q.Encode()
	}
	result, err := decode[[]Template](c, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func (c *Client) GetTemplate(id int) (*Template, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/templates/%d", id), nil)
	if err != nil {
		return nil, err
	}
	result, err := decode[Template](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) CreateTemplate(t NewTemplate) (*Template, error) {
	req, err := c.newRequest(http.MethodPost, "/templates", t)
	if err != nil {
		return nil, err
	}
	result, err := decode[Template](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) UpdateTemplate(id int, t UpdateTemplate) (*Template, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("/templates/%d", id), t)
	if err != nil {
		return nil, err
	}
	result, err := decode[Template](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) SetDefaultTemplate(id int) error {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("/templates/%d/default", id), nil)
	if err != nil {
		return err
	}
	_, err = decode[bool](c, req)
	return err
}


func (c *Client) DeleteTemplate(id int) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/templates/%d", id), nil)
	if err != nil {
		return err
	}
	_, err = decode[bool](c, req)
	return err
}


func (c *Client) SendTransactional(msg TransactionalMessage) error {
	req, err := c.newRequest(http.MethodPost, "/tx", msg)
	if err != nil {
		return err
	}
	_, err = decode[bool](c, req)
	return err
}


func (c *Client) GetBounces(params GetBouncesParams) (*BounceList, error) {
	req, err := c.newRequest(http.MethodGet, "/bounces", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	setPage(q, params.Page, params.PerPage)
	if params.CampaignID > 0 {
		q.Set("campaign_id", strconv.Itoa(params.CampaignID))
	}
	if params.Source != "" {
		q.Set("source", params.Source)
	}
	if params.OrderBy != "" {
		q.Set("order_by", params.OrderBy)
	}
	if params.Order != "" {
		q.Set("order", params.Order)
	}
	req.URL.RawQuery = q.Encode()
	result, err := decode[BounceList](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) DeleteBounce(id int) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/bounces/%d", id), nil)
	if err != nil {
		return err
	}
	_, err = decode[bool](c, req)
	return err
}


func (c *Client) DeleteBounces(ids []int, all bool) error {
	req, err := c.newRequest(http.MethodDelete, "/bounces", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	if all {
		q.Set("all", "true")
	}
	for _, id := range ids {
		q.Add("id", strconv.Itoa(id))
	}
	req.URL.RawQuery = q.Encode()
	_, err = decode[bool](c, req)
	return err
}

func (c *Client) GetMedia() ([]MediaFileObject, error) {
	req, err := c.newRequest(http.MethodGet, "/media", nil)
	if err != nil {
		return nil, err
	}
	result, err := decode[[]MediaFileObject](c, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func (c *Client) DeleteMedia(id int) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/media/%d", id), nil)
	if err != nil {
		return err
	}
	_, err = decode[bool](c, req)
	return err
}

func (c *Client) GetSettings() (*Settings, error) {
	req, err := c.newRequest(http.MethodGet, "/settings", nil)
	if err != nil {
		return nil, err
	}
	result, err := decode[Settings](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) UpdateSettings(settings Settings) (*Settings, error) {
	req, err := c.newRequest(http.MethodPut, "/settings", settings)
	if err != nil {
		return nil, err
	}
	result, err := decode[Settings](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) GetImportStatus() (*ImportStatus, error) {
	req, err := c.newRequest(http.MethodGet, "/import/subscribers", nil)
	if err != nil {
		return nil, err
	}
	result, err := decode[ImportStatus](c, req)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
