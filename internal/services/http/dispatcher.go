package http_service

import util "github.com/angeloevangelista/go-bucket/internal/utils"

type HttpRequestDispatcher[T any] struct {
	defaultHeaders map[string]string
	queryParams    map[string]string

	Url          string
	ResponseData *string
	RequestData  *string
	StatusCode   *int
}

type HttpRequestDispatcherCreateOptions struct {
	Url         string
	Headers     map[string]string
	QueryParams map[string]string
	RequestData *string
}

func (httpRequestDispatcher *HttpRequestDispatcher[T]) DeserializeResponse() (
	*T, error,
) {
	return util.DeserializeObject[T](*httpRequestDispatcher.ResponseData)
}

func CreateDispatcher[T any](
	httpClientCreateOptions HttpRequestDispatcherCreateOptions,
) HttpRequestDispatcher[T] {
	return HttpRequestDispatcher[T]{
		defaultHeaders: httpClientCreateOptions.Headers,
		queryParams:    httpClientCreateOptions.QueryParams,

		Url:         httpClientCreateOptions.Url,
		RequestData: httpClientCreateOptions.RequestData,
	}
}
