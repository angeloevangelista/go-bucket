package bitbucket

import (
	"fmt"
	"strings"

	bitbucket_models "github.com/angeloevangelista/go-bucket/internal/bitbucket/models"
	http_service "github.com/angeloevangelista/go-bucket/internal/services/http"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func (bitbucketClient *BitbucketClient) PipelinesHandler() *PipelinesHandler {
	return &PipelinesHandler{
		client: *bitbucketClient,
	}
}

type PipelinesHandler struct {
	client BitbucketClient
}

func (PipelinesHandler *PipelinesHandler) ListByRepository(
	workspaceCode string,
	repositoryCode string,
	paginationOptions bitbucket_models.PaginationOptions,
) (
	pipelinesResponse *bitbucket_models.PaginatedResponse[bitbucket_models.PipelineResponse],
	err error,
) {
	httpDispatcher := http_service.CreateDispatcher[bitbucket_models.PaginatedResponse[bitbucket_models.PipelineResponse]](
		http_service.HttpRequestDispatcherCreateOptions{
			Url: strings.Join(
				[]string{
					"https://api.bitbucket.org/2.0/repositories",
					workspaceCode,
					repositoryCode,
					"pipelines/",
				},
				"/",
			),
			Headers: map[string]string{
				"Authorization": "Bearer " + PipelinesHandler.client.accessToken,
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

	pipelinesResponse, err = httpDispatcher.DeserializeResponse()

	if util.CheckError(err) {
		return nil, err
	}

	return pipelinesResponse, nil
}
