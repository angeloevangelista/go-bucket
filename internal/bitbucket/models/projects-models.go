package bitbucket_models

type ProjectResponse struct {
	Type                    string    `json:"type"`
	Owner                   Owner     `json:"owner"`
	Workspace               Workspace `json:"workspace"`
	Key                     string    `json:"key"`
	UUID                    string    `json:"uuid"`
	IsPrivate               bool      `json:"is_private"`
	Name                    string    `json:"name"`
	Description             string    `json:"description"`
	Links                   Links     `json:"links"`
	CreatedOn               string    `json:"created_on"`
	UpdatedOn               string    `json:"updated_on"`
	HasPubliclyVisibleRepos bool      `json:"has_publicly_visible_repos"`
}

type Links struct {
	Self         Avatar  `json:"self"`
	HTML         Avatar  `json:"html"`
	Repositories *Avatar `json:"repositories,omitempty"`
	Avatar       Avatar  `json:"avatar"`
}

type Workspace struct {
	Type  string `json:"type"`
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Links Links  `json:"links"`
}
