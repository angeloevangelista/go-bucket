package main

import (
	"github.com/angeloevangelista/go-bucket/internal/bitbucket"
	bitbucket_models "github.com/angeloevangelista/go-bucket/internal/bitbucket/models"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func panicIfError(err error) {
	if util.CheckError(err) {
		panic(err)
	}
}

func listWorkspaces(
	bitbucketClient *bitbucket.BitbucketClient,
) []bitbucket_models.WorkspaceResponse {
	var workspaces []bitbucket_models.WorkspaceResponse

	hasMorePages := true

	for hasMorePages {
		workspacesResponse, err := bitbucketClient.WorkspacesHandler().ListForCurrentUser(
			bitbucket_models.PaginationOptions{
				PageLimit:  100,
				PageNumber: 1,
			},
		)

		panicIfError(err)

		hasMorePages = workspacesResponse.Next != ""
		workspaces = append(workspaces, workspacesResponse.Values...)
	}

	return workspaces
}

func listRepositories(
	bitbucketClient *bitbucket.BitbucketClient,
	workspaceSlug string,
) []bitbucket_models.RepositoryResponse {
	var repositories []bitbucket_models.RepositoryResponse

	hasMorePages := true

	for hasMorePages {
		repositoriesResponse, err := bitbucketClient.RepositoriesHandler().ListByWorkspace(
			workspaceSlug,
			bitbucket_models.PaginationOptions{
				PageLimit:  100,
				PageNumber: 1,
			},
		)

		panicIfError(err)

		hasMorePages = repositoriesResponse.Next != ""
		repositories = append(repositories, repositoriesResponse.Values...)
	}

	return repositories
}

func listCommits(
	bitbucketClient *bitbucket.BitbucketClient,
	workspaceSlug string,
	repositorySlug string,
) []bitbucket_models.CommitResponse {
	var commits []bitbucket_models.CommitResponse

	hasMorePages := true

	for hasMorePages {
		commitsResponse, err := bitbucketClient.CommitsHandler().ListByRepository(
			workspaceSlug,
			repositorySlug,
			bitbucket_models.PaginationOptions{
				PageLimit:  100,
				PageNumber: 1,
			},
		)

		panicIfError(err)

		hasMorePages = commitsResponse.Next != ""
		commits = append(commits, commitsResponse.Values...)
	}

	return commits
}
