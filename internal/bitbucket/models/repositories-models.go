package bitbucket_models

type RepositoryResponse struct {
	Type             string                  `json:"type"`
	FullName         string                  `json:"full_name"`
	Links            RepositoryResponseLinks `json:"links"`
	Name             string                  `json:"name"`
	Slug             string                  `json:"slug"`
	Description      string                  `json:"description"`
	SCM              string                  `json:"scm"`
	Website          interface{}             `json:"website"`
	Owner            Owner                   `json:"owner"`
	Workspace        Project                 `json:"workspace"`
	IsPrivate        bool                    `json:"is_private"`
	Project          Project                 `json:"project"`
	ForkPolicy       string                  `json:"fork_policy"`
	CreatedOn        string                  `json:"created_on"`
	UpdatedOn        string                  `json:"updated_on"`
	Size             int64                   `json:"size"`
	Language         string                  `json:"language"`
	HasIssues        bool                    `json:"has_issues"`
	HasWiki          bool                    `json:"has_wiki"`
	UUID             string                  `json:"uuid"`
	Mainbranch       Mainbranch              `json:"mainbranch"`
	OverrideSettings OverrideSettings        `json:"override_settings"`
}

type RepositoryResponseLinks struct {
	Self         Avatar  `json:"self"`
	HTML         Avatar  `json:"html"`
	Avatar       Avatar  `json:"avatar"`
	Pullrequests Avatar  `json:"pullrequests"`
	Commits      Avatar  `json:"commits"`
	Forks        Avatar  `json:"forks"`
	Watchers     Avatar  `json:"watchers"`
	Branches     Avatar  `json:"branches"`
	Tags         Avatar  `json:"tags"`
	Downloads    Avatar  `json:"downloads"`
	Source       Avatar  `json:"source"`
	Clone        []Clone `json:"clone"`
	Hooks        Avatar  `json:"hooks"`
}

type Clone struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

type Mainbranch struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type OverrideSettings struct {
	DefaultMergeStrategy bool `json:"default_merge_strategy"`
	BranchingModel       bool `json:"branching_model"`
}

type OwnerLinks struct {
	Self   Avatar `json:"self"`
	Avatar Avatar `json:"avatar"`
	HTML   Avatar `json:"html"`
}

type Project struct {
	Type  string     `json:"type"`
	Key   *string    `json:"key,omitempty"`
	UUID  string     `json:"uuid"`
	Name  string     `json:"name"`
	Links OwnerLinks `json:"links"`
	Slug  *string    `json:"slug,omitempty"`
}
