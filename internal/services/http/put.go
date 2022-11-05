package http_service

import (
	"bytes"
	"io/ioutil"
	"net/http"

	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func (httpRequestDispatcher *HttpRequestDispatcher[T]) Put() error {
	var requestBodyBytes []byte

	if httpRequestDispatcher.RequestData != nil {
		requestBodyBytes = []byte(*httpRequestDispatcher.RequestData)
	}

	nativeRequest, err := http.NewRequest(
		"PUT",
		httpRequestDispatcher.Url,
		bytes.NewBuffer(requestBodyBytes),
	)

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
