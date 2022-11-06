package bitbucket_models

type UserResponse struct {
	Type      string            `json:"type"`
	User      User              `json:"user"`
	Workspace Workspace         `json:"workspace"`
	Links     UserResponseLinks `json:"links"`
}

type UserResponseLinks struct {
	Self Self `json:"self"`
}

type Self struct {
	Href string `json:"href"`
}

type User struct {
	DisplayName string    `json:"display_name"`
	Links       UserLinks `json:"links"`
	Type        string    `json:"type"`
	UUID        string    `json:"uuid"`
	AccountID   string    `json:"account_id"`
	Nickname    string    `json:"nickname"`
}

type UserLinks struct {
	Self   Self `json:"self"`
	Avatar Self `json:"avatar"`
	HTML   Self `json:"html"`
}
