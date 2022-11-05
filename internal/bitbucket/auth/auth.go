package bitbucket

import (
	"net/url"

	http_service "github.com/angeloevangelista/go-bucket/internal/services/http"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func GetAccessToken(getAccessTokenOptions GetAccessTokenOptions) (*string, error) {
	var err error

	formData := url.Values{
		"client_id":     {getAccessTokenOptions.ClientId},
		"client_secret": {getAccessTokenOptions.ClientSecret},
		"grant_type":    {"client_credentials"},
	}.Encode()

	httpDispatcher := http_service.CreateDispatcher[BitbucketAccessTokenResponse](
		http_service.HttpRequestDispatcherCreateOptions{
			Url: "https://bitbucket.org/site/oauth2/access_token",
			Headers: map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			},
			RequestData: &formData,
		},
	)

	err = httpDispatcher.Post()

	if util.CheckError(err) {
		return nil, err
	}

	response, err := httpDispatcher.DeserializeResponse()

	if util.CheckError(err) {
		return nil, err
	}

	return &response.AccessToken, nil
}
