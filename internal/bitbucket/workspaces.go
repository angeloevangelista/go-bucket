package bitbucket

import (
	"fmt"

	bitbucket_models "github.com/angeloevangelista/go-bucket/internal/bitbucket/models"
	http_service "github.com/angeloevangelista/go-bucket/internal/services/http"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func (bitbucketClient *BitbucketClient) WorkspacesHandler() *WorkspacesHandler {
	return &WorkspacesHandler{
		client: *bitbucketClient,
	}
}

type WorkspacesHandler struct {
	client BitbucketClient
}

func (workspacesHandler *WorkspacesHandler) ListForCurrentUser(
	paginationOptions bitbucket_models.PaginationOptions,
) (
	workspacesResponse *bitbucket_models.PaginatedResponse[bitbucket_models.WorkspaceResponse],
	err error,
) {
	httpDispatcher := http_service.CreateDispatcher[bitbucket_models.PaginatedResponse[bitbucket_models.WorkspaceResponse]](
		http_service.HttpRequestDispatcherCreateOptions{
			Url: "https://api.bitbucket.org/2.0/workspaces",
			Headers: map[string]string{
				"Authorization": "Bearer " + workspacesHandler.client.accessToken,
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

	workspacesResponse, err = httpDispatcher.DeserializeResponse()

	if util.CheckError(err) {
		return nil, err
	}

	return workspacesResponse, nil
}
