package main

import (
	"log"
	"strings"

	"github.com/angeloevangelista/go-bucket/internal/bitbucket"
	bitbucket_models "github.com/angeloevangelista/go-bucket/internal/bitbucket/models"
	secrets_storage_service "github.com/angeloevangelista/go-bucket/internal/services/secrets-storage"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func main() {
	bitbucketEvents := []BitbucketMappedEvent{}

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

	workspaces := listWorkspaces(bitbucketClient)

	for _, workspace := range workspaces {
		repositories := listRepositories(bitbucketClient, workspace.Slug)

		for _, repository := range repositories {
			commits := listCommits(bitbucketClient, workspace.Slug, repository.Slug)

			bitbucketEvent := BitbucketMappedEvent{
				UserDocuments:   []User{},
				CommitDocuments: []Commit{},
				RepositoryDocument: Repository{
					BitbucketID: util.SanitizeBitbucketUUID(repository.UUID),
					Name:        repository.Name,
					FullName:    repository.FullName,
					Private:     repository.IsPrivate,
				},
				ProjectDocument: Project{
					BitbucketID: util.SanitizeBitbucketUUID(repository.Project.UUID),
					Name:        repository.Project.Name,
					Key:         *repository.Project.Key,
				},
			}

			for _, commit := range commits {
				authorEmail := strings.Replace(
					strings.Split(commit.Author.Raw, "\u003c")[1],
					"\u003e",
					"",
					1,
				)

				userIsAlreadyInList := false

				for _, user := range bitbucketEvent.UserDocuments {
					if user.Email == authorEmail {
						userIsAlreadyInList = true
						break
					}
				}

				if !userIsAlreadyInList {
					bitbucketEvent.UserDocuments = append(
						bitbucketEvent.UserDocuments,
						User{
							Name:      commit.Author.User.DisplayName,
							Email:     authorEmail,
							AccountId: commit.Author.User.AccountID,
						},
					)
				}

				bitbucketEvent.CommitDocuments = append(
					bitbucketEvent.CommitDocuments,
					Commit{
						Hash:        commit.Hash,
						Date:        commit.Date,
						Message:     commit.Message,
						AuthorEmail: authorEmail,
						AuthorName:  commit.Author.User.DisplayName,
					},
				)
			}

			bitbucketEvents = append(bitbucketEvents, bitbucketEvent)
		}
	}

	json, err := util.SerializeObject(bitbucketEvents)

	panicIfError(err)

	log.Printf(*json)
}
