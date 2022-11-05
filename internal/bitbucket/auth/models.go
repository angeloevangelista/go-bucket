package bitbucket

type GetAccessTokenOptions struct {
	ClientId     string
	ClientSecret string
}

type BitbucketAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	Scopes       string `json:"scopes"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	State        string `json:"state"`
	RefreshToken string `json:"refresh_token"`
}
