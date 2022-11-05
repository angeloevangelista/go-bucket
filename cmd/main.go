package main

import (
	"log"

	bitbucket "github.com/angeloevangelista/go-bucket/internal/bitbucket/auth"
	util "github.com/angeloevangelista/go-bucket/internal/utils"
)

func main() {
	log.Print("🚀 It works!")

	accessToken, err := bitbucket.GetAccessToken(bitbucket.GetAccessTokenOptions{
		ClientId:     "Ooops",
		ClientSecret: "Someone forgot to remove this",
	})

	if util.CheckError(err) {
		panic(err)
	}

	log.Print(*accessToken)
}
