package bitbucket

import (
	"fmt"

	bitbucket_models "github.com/angeloevangelista/go-bucket/internal/bitbucket/models"
	http_service "github.com/angeloevangelista/go-bucket/internal/services/http"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func (bitbucketClient *BitbucketClient) RepositoriesHandler() *RepositoriesHandler {
	return &RepositoriesHandler{
		client: *bitbucketClient,
	}
}

type RepositoriesHandler struct {
	client BitbucketClient
}

// It lists all public repositories over the Bitbucket API, not only yours.
func (RepositoriesHandler *RepositoriesHandler) ListPublic(
	paginationOptions bitbucket_models.PaginationOptions,
) (
	repositoriesResponse *bitbucket_models.PaginatedResponse[bitbucket_models.RepositoryResponse],
	err error,
) {
	httpDispatcher := http_service.CreateDispatcher[bitbucket_models.PaginatedResponse[bitbucket_models.RepositoryResponse]](
		http_service.HttpRequestDispatcherCreateOptions{
			Url: "https://api.bitbucket.org/2.0/repositories",
			Headers: map[string]string{
				"Authorization": "Bearer " + RepositoriesHandler.client.accessToken,
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

	repositoriesResponse, err = httpDispatcher.DeserializeResponse()

	if util.CheckError(err) {
		return nil, err
	}

	return repositoriesResponse, nil
}

// It lists all repositories for the current user and workspace.
func (RepositoriesHandler *RepositoriesHandler) ListByWorkspace(
	workspaceCode string,
	paginationOptions bitbucket_models.PaginationOptions,
) (
	repositoriesResponse *bitbucket_models.PaginatedResponse[bitbucket_models.RepositoryResponse],
	err error,
) {
	httpDispatcher := http_service.CreateDispatcher[bitbucket_models.PaginatedResponse[bitbucket_models.RepositoryResponse]](
		http_service.HttpRequestDispatcherCreateOptions{
			Url: "https://api.bitbucket.org/2.0/repositories/" + workspaceCode,
			Headers: map[string]string{
				"Authorization": "Bearer " + RepositoriesHandler.client.accessToken,
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

	repositoriesResponse, err = httpDispatcher.DeserializeResponse()

	if util.CheckError(err) {
		return nil, err
	}

	return repositoriesResponse, nil
}
