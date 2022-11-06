package bitbucket_models

type CommitResponse struct {
	Type       string              `json:"type"`
	Hash       string              `json:"hash"`
	Date       string              `json:"date"`
	Author     Author              `json:"author"`
	Message    string              `json:"message"`
	Summary    Summary             `json:"summary"`
	Links      CommitResponseLinks `json:"links"`
	Parents    []interface{}       `json:"parents"`
	Rendered   Rendered            `json:"rendered"`
	Repository Repository          `json:"repository"`
}

type Author struct {
	Type string `json:"type"`
	Raw  string `json:"raw"`
	User User   `json:"user"`
}

type Approve struct {
	Href string `json:"href"`
}

type CommitResponseLinks struct {
	Self     Approve `json:"self"`
	HTML     Approve `json:"html"`
	Diff     Approve `json:"diff"`
	Approve  Approve `json:"approve"`
	Comments Approve `json:"comments"`
	Statuses Approve `json:"statuses"`
	Patch    Approve `json:"patch"`
}

type Rendered struct {
	Message Summary `json:"message"`
}

type Summary struct {
	Type   string `json:"type"`
	Raw    string `json:"raw"`
	Markup string `json:"markup"`
	HTML   string `json:"html"`
}

type Repository struct {
	Type     string    `json:"type"`
	FullName string    `json:"full_name"`
	Links    UserLinks `json:"links"`
	Name     string    `json:"name"`
	UUID     string    `json:"uuid"`
}
