package bitbucket_models

type WorkspaceResponse struct {
	Type      string `json:"type"`
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	IsPrivate bool   `json:"is_private"`
	CreatedOn string `json:"created_on"`
}
