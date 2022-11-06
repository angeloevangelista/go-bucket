package main

import (
	"log"

	"github.com/angeloevangelista/go-bucket/internal/bitbucket"
	bitbucket_models "github.com/angeloevangelista/go-bucket/internal/bitbucket/models"
	secrets_storage_service "github.com/angeloevangelista/go-bucket/internal/services/secrets-storage"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func panicIfError(err error) {
	if util.CheckError(err) {
		panic(err)
	}
}

func main() {
	log.Print("🚀 It works!")

	clientId, err := secrets_storage_service.GetSecret(
		"BITBUCKET_LEGACY_CLIENT_ID",
	)

	panicIfError(err)

	clientSecret, err := secrets_storage_service.GetSecret(
		"BITBUCKET_LEGACY_CLIENT_SECRET",
	)

	panicIfError(err)

	bitbucketClient, err := bitbucket.GetClient(
		bitbucket_models.GetBitbucketClientOptions{
			ClientId:     *clientId,
			ClientSecret: *clientSecret,
		},
	)

	panicIfError(err)

	workspacesResponse, err := bitbucketClient.WorkspacesHandler().ListForCurrentUser(
		bitbucket_models.PaginationOptions{
			PageLimit:  100,
			PageNumber: 1,
		},
	)

	panicIfError(err)

	for _, workspace := range workspacesResponse.Values {
		log.Printf("Workspace: %s", workspace.Slug)

		repositoriesResponse, err := bitbucketClient.RepositoriesHandler().ListByWorkspace(
			workspace.Slug,
			bitbucket_models.PaginationOptions{
				PageLimit:  100,
				PageNumber: 1,
			},
		)

		panicIfError(err)

		for _, repository := range repositoriesResponse.Values {
			log.Printf("Repository: %s", repository.Name)
		}

		projectsResponse, err := bitbucketClient.ProjectsHandler().ListByWorkspace(
			workspace.Slug,
			bitbucket_models.PaginationOptions{
				PageLimit:  100,
				PageNumber: 1,
			},
		)

		panicIfError(err)

		for _, project := range projectsResponse.Values {
			log.Printf("Project: %s", project.Name)
		}

		log.Printf("-----")
		log.Printf("")
	}
}
