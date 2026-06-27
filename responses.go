package listmonk

import "fmt"

type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("listmonk: %d %s", e.Code, e.Message)
}

type Page[T any] struct {
	Results []T    `json:"results"`
	Total   int    `json:"total"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Query   string `json:"query,omitempty"`
}


type BounceList struct {
	Results []BounceRecord `json:"results"`
	Total   int            `json:"total"`
	Page    int            `json:"page"`
	PerPage int            `json:"per_page"`
}


type DashboardCounts struct {
	Subscribers *DashboardSubscriberCounts `json:"subscribers,omitempty"`
	Lists       *DashboardListCounts       `json:"lists,omitempty"`
	Campaigns   *DashboardCampaignCounts   `json:"campaigns,omitempty"`
	Messages    int                        `json:"messages,omitempty"`
}

type DashboardSubscriberCounts struct {
	Total       int `json:"total,omitempty"`
	Blocklisted int `json:"blocklisted,omitempty"`
	Orphans     int `json:"orphans,omitempty"`
}

type DashboardListCounts struct {
	Total       int `json:"total,omitempty"`
	Public      int `json:"public,omitempty"`
	Private     int `json:"private,omitempty"`
	OptinSingle int `json:"optin_single,omitempty"`
	OptinDouble int `json:"optin_double,omitempty"`
}

type DashboardCampaignCounts struct {
	Total    int                        `json:"total,omitempty"`
	ByStatus *DashboardCampaignByStatus `json:"by_status,omitempty"`
}

type DashboardCampaignByStatus struct {
	Draft int `json:"draft,omitempty"`
}

type DashboardCharts struct {
	CampaignViews []DashboardChartPoint `json:"campaign_views,omitempty"`
	LinkClicks    []DashboardChartPoint `json:"link_clicks,omitempty"`
}

type DashboardChartPoint struct {
	Date  string `json:"date,omitempty"`
	Count int    `json:"count,omitempty"`
}



type ServerConfig struct {
	Version      string             `json:"version,omitempty"`
	Update       string             `json:"update,omitempty"`
	Lang         string             `json:"lang,omitempty"`
	Langs        []ServerConfigLang `json:"langs,omitempty"`
	Messengers   []string           `json:"messengers,omitempty"`
	NeedsRestart bool               `json:"needs_restart,omitempty"`
}

type ServerConfigLang struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}



type ImportStatus struct {
	Name     string `json:"name,omitempty"`
	Status   string `json:"status,omitempty"`
	Total    int    `json:"total,omitempty"`
	Imported int    `json:"imported,omitempty"`
}
