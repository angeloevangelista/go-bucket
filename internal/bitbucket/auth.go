package bitbucket

import (
	"net/url"

	bitbucket_models "github.com/angeloevangelista/go-bucket/internal/bitbucket/models"
	http_service "github.com/angeloevangelista/go-bucket/internal/services/http"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func GetAccessToken(getAccessTokenOptions bitbucket_models.GetAccessTokenOptions) (
	accessToken *string,
	refreshToken *string,
	err error,
) {
	formData := url.Values{
		"client_id":     {getAccessTokenOptions.ClientId},
		"client_secret": {getAccessTokenOptions.ClientSecret},
		"grant_type":    {"client_credentials"},
	}.Encode()

	httpDispatcher := http_service.CreateDispatcher[bitbucket_models.AccessTokenResponse](
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
		return nil, nil, err
	}

	response, err := httpDispatcher.DeserializeResponse()

	if util.CheckError(err) {
		return nil, nil, err
	}

	return &response.AccessToken, &response.RefreshToken, nil
}

func (bitbucketClient *BitbucketClient) RefreshAuth() (err error) {
	formData := url.Values{
		"client_id":     {bitbucketClient.clientId},
		"client_secret": {bitbucketClient.clientSecret},
		"refresh_token": {bitbucketClient.refreshToken},
		"grant_type":    {"refresh_token"},
	}.Encode()

	httpDispatcher := http_service.CreateDispatcher[bitbucket_models.AccessTokenResponse](
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
		return err
	}

	response, err := httpDispatcher.DeserializeResponse()

	if util.CheckError(err) {
		return err
	}

	bitbucketClient.accessToken = response.AccessToken
	bitbucketClient.refreshToken = response.RefreshToken

	return nil
}
