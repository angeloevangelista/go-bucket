package bitbucket

import (
	bitbucket_models "github.com/angeloevangelista/go-bucket/internal/bitbucket/models"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

type BitbucketClient struct {
	clientId     string
	clientSecret string
	accessToken  string
	refreshToken string
}

func GetClient(
	getBitbucketClientOptions bitbucket_models.GetBitbucketClientOptions,
) (bitbucketClient *BitbucketClient, err error) {
	accessToken, refreshToken, err := GetAccessToken(
		bitbucket_models.GetAccessTokenOptions{
			ClientId:     getBitbucketClientOptions.ClientId,
			ClientSecret: getBitbucketClientOptions.ClientSecret,
		},
	)

	if util.CheckError(err) {
		return nil, err
	}

	bitbucketClient = &BitbucketClient{
		clientId:     getBitbucketClientOptions.ClientId,
		clientSecret: getBitbucketClientOptions.ClientSecret,
		accessToken:  *accessToken,
		refreshToken: *refreshToken,
	}

	return bitbucketClient, nil
}
