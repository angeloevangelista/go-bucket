package bitbucket_models

type PaginatedResponse[T any] struct {
	Values   []T    `json:"values"`
	Pagelen  int64  `json:"pagelen"`
	Size     int64  `json:"size"`
	Page     int64  `json:"page"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type PaginationOptions struct {
	PageLimit  int64
	PageNumber int64
}

type Avatar struct {
	Href string `json:"href"`
}

type Owner struct {
	DisplayName string `json:"display_name"`
	Links       Links  `json:"links"`
	Type        string `json:"type"`
	UUID        string `json:"uuid"`
	AccountID   string `json:"account_id"`
	Nickname    string `json:"nickname"`
}
