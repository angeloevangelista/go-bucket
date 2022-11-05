package main

import (
	"log"

	bitbucket "github.com/angeloevangelista/go-bucket/internal/bitbucket/auth"
	secrets_storage_service "github.com/angeloevangelista/go-bucket/internal/services/secrets-storage"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func main() {
	log.Print("🚀 It works!")

	clientId, err := secrets_storage_service.GetSecret("BITBUCKET_LEGACY_CLIENT_ID")

	if util.CheckError(err) {
		panic(err)
	}

	clientSecret, err := secrets_storage_service.GetSecret("BITBUCKET_LEGACY_CLIENT_SECRET")

	if util.CheckError(err) {
		panic(err)
	}

	accessToken, err := bitbucket.GetAccessToken(bitbucket.GetAccessTokenOptions{
		ClientId:     *clientId,
		ClientSecret: *clientSecret,
	})

	if util.CheckError(err) {
		panic(err)
	}

	log.Print(*accessToken)
}
