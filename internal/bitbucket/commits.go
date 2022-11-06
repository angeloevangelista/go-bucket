package bitbucket

import (
	"fmt"
	"strings"

	bitbucket_models "github.com/angeloevangelista/go-bucket/internal/bitbucket/models"
	http_service "github.com/angeloevangelista/go-bucket/internal/services/http"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func (bitbucketClient *BitbucketClient) CommitsHandler() *CommitsHandler {
	return &CommitsHandler{
		client: *bitbucketClient,
	}
}

type CommitsHandler struct {
	client BitbucketClient
}

func (commitsHandler *CommitsHandler) ListByRepository(
	workspaceCode string,
	repositoryCode string,
	paginationOptions bitbucket_models.PaginationOptions,
) (
	projectsResponse *bitbucket_models.PaginatedResponse[bitbucket_models.CommitResponse],
	err error,
) {
	httpDispatcher := http_service.CreateDispatcher[bitbucket_models.PaginatedResponse[bitbucket_models.CommitResponse]](
		http_service.HttpRequestDispatcherCreateOptions{
			Url: strings.Join(
				[]string{
					"https://api.bitbucket.org/2.0/repositories",
					workspaceCode,
					repositoryCode,
					"commits",
				},
				"/",
			),
			Headers: map[string]string{
				"Authorization": "Bearer " + commitsHandler.client.accessToken,
			},
			QueryParams: map[string]string{
				"pagelen": fmt.Sprint(paginationOptions.PageLimit),
				"page":    fmt.Sprint(paginationOptions.PageNumber),
			},
		},
	)

	err = httpDispatcher.Get()

	if util.CheckError(err) {
		return nil, err
	}

	projectsResponse, err = httpDispatcher.DeserializeResponse()

	if util.CheckError(err) {
		return nil, err
	}

	return projectsResponse, nil
}
