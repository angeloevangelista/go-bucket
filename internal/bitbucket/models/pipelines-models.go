package bitbucket_models

type PipelineResponse struct {
	Type              string                `json:"type"`
	UUID              string                `json:"uuid"`
	Repository        Repository            `json:"repository"`
	State             State                 `json:"state"`
	BuildNumber       int64                 `json:"build_number"`
	Creator           Creator               `json:"creator"`
	CreatedOn         string                `json:"created_on"`
	Target            Target                `json:"target"`
	Trigger           Trigger               `json:"trigger"`
	RunNumber         int64                 `json:"run_number"`
	DurationInSeconds int64                 `json:"duration_in_seconds"`
	BuildSecondsUsed  int64                 `json:"build_seconds_used"`
	FirstSuccessful   bool                  `json:"first_successful"`
	Expired           bool                  `json:"expired"`
	Links             PipelineResponseLinks `json:"links"`
	HasVariables      bool                  `json:"has_variables"`
}

type Creator struct {
	DisplayName string       `json:"display_name"`
	Links       CreatorLinks `json:"links"`
	Type        string       `json:"type"`
	UUID        string       `json:"uuid"`
	AccountID   string       `json:"account_id"`
	Nickname    string       `json:"nickname"`
}

type CreatorLinks struct {
	Self   Self `json:"self"`
	HTML   Self `json:"html"`
	Avatar Self `json:"avatar"`
}

type PipelineResponseLinks struct {
	Self  Self `json:"self"`
	Steps Self `json:"steps"`
}

type State struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Stage Trigger `json:"stage"`
}

type Trigger struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Target struct {
	Type     string   `json:"type"`
	RefType  string   `json:"ref_type"`
	RefName  string   `json:"ref_name"`
	Selector Selector `json:"selector"`
	Commit   Commit   `json:"commit"`
}

type Commit struct {
	Type  string      `json:"type"`
	Hash  string      `json:"hash"`
	Links CommitLinks `json:"links"`
}

type CommitLinks struct {
	Self Self `json:"self"`
	HTML Self `json:"html"`
}

type Selector struct {
	Type string `json:"type"`
}
