package bitbucket

import (
	"fmt"

	bitbucket_models "github.com/angeloevangelista/go-bucket/internal/bitbucket/models"
	http_service "github.com/angeloevangelista/go-bucket/internal/services/http"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func (bitbucketClient *BitbucketClient) ProjectsHandler() *ProjectsHandler {
	return &ProjectsHandler{
		client: *bitbucketClient,
	}
}

type ProjectsHandler struct {
	client BitbucketClient
}

func (projectsHandler *ProjectsHandler) ListByWorkspace(
	workspaceCode string,
	paginationOptions bitbucket_models.PaginationOptions,
) (
	projectsResponse *bitbucket_models.PaginatedResponse[bitbucket_models.ProjectResponse],
	err error,
) {
	httpDispatcher := http_service.CreateDispatcher[bitbucket_models.PaginatedResponse[bitbucket_models.ProjectResponse]](
		http_service.HttpRequestDispatcherCreateOptions{
			Url: "https://api.bitbucket.org/2.0/workspaces/" + workspaceCode + "/projects",
			Headers: map[string]string{
				"Authorization": "Bearer " + projectsHandler.client.accessToken,
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
