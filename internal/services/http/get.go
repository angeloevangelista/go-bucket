package http_service

import (
	"io/ioutil"
	"net/http"

	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func (httpRequestDispatcher *HttpRequestDispatcher[T]) Get() error {
	nativeRequest, err := http.NewRequest("GET", httpRequestDispatcher.Url, nil)

	if util.CheckError(err) {
		return nil
	}

	for key, value := range httpRequestDispatcher.defaultHeaders {
		nativeRequest.Header.Add(key, value)
	}

	queryParams := nativeRequest.URL.Query()

	for key, value := range httpRequestDispatcher.queryParams {
		queryParams.Add(key, value)
	}

	nativeRequest.URL.RawQuery = queryParams.Encode()

	nativeHttpClient := &http.Client{}
	nativeHttpResponse, err := nativeHttpClient.Do(nativeRequest)

	if util.CheckError(err) {
		return nil
	}

	httpRequestDispatcher.StatusCode = &nativeHttpResponse.StatusCode

	bodyBytes, err := ioutil.ReadAll(nativeHttpResponse.Body)

	if util.CheckError(err) {
		return nil
	}

	rawBody := string(bodyBytes)

	if util.CheckError(err) {
		return nil
	}

	httpRequestDispatcher.ResponseData = &rawBody

	return nil
}
