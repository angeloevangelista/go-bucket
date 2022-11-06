package bitbucket

import (
	"fmt"

	bitbucket_models "github.com/angeloevangelista/go-bucket/internal/bitbucket/models"
	http_service "github.com/angeloevangelista/go-bucket/internal/services/http"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func (bitbucketClient *BitbucketClient) UsersHandler() *UsersHandler {
	return &UsersHandler{
		client: *bitbucketClient,
	}
}

type UsersHandler struct {
	client BitbucketClient
}

func (usersHandler *UsersHandler) ListByWorkspace(
	workspaceCode string,
	paginationOptions bitbucket_models.PaginationOptions,
) (
	projectsResponse *bitbucket_models.PaginatedResponse[bitbucket_models.UserResponse],
	err error,
) {
	httpDispatcher := http_service.CreateDispatcher[bitbucket_models.PaginatedResponse[bitbucket_models.UserResponse]](
		http_service.HttpRequestDispatcherCreateOptions{
			Url: "https://api.bitbucket.org/2.0/workspaces/" + workspaceCode + "/members",
			Headers: map[string]string{
				"Authorization": "Bearer " + usersHandler.client.accessToken,
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
