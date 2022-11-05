package bitbucket_models

type PaginatedResponse[T any] struct {
	Values  []T   `json:"values"`
	Pagelen int64 `json:"pagelen"`
	Size    int64 `json:"size"`
	Page    int64 `json:"page"`
}

type PaginationOptions struct {
	PageLimit  int64
	PageNumber int64
}
