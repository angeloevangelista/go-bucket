package main

import (
	"log"

	"github.com/angeloevangelista/go-bucket/internal/bitbucket"
	bitbucket_models "github.com/angeloevangelista/go-bucket/internal/bitbucket/models"
	secrets_storage_service "github.com/angeloevangelista/go-bucket/internal/services/secrets-storage"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func main() {
	log.Print("🚀 It works!")

	clientId, err := secrets_storage_service.GetSecret(
		"BITBUCKET_LEGACY_CLIENT_ID",
	)

	if util.CheckError(err) {
		panic(err)
	}

	clientSecret, err := secrets_storage_service.GetSecret(
		"BITBUCKET_LEGACY_CLIENT_SECRET",
	)

	if util.CheckError(err) {
		panic(err)
	}

	bitbucketClient, err := bitbucket.GetClient(
		bitbucket_models.GetBitbucketClientOptions{
			ClientId:     *clientId,
			ClientSecret: *clientSecret,
		},
	)

	if util.CheckError(err) {
		panic(err)
	}

	workspacesResponse, err := bitbucketClient.WorkspacesHandler().ListForCurrentUser(
		bitbucket_models.PaginationOptions{
			PageLimit:  10,
			PageNumber: 1,
		},
	)

	if util.CheckError(err) {
		panic(err)
	}

	for i, workspace := range workspacesResponse.Values {
		log.Printf("Workspace %d: %s", i, workspace.Name)
	}
}
