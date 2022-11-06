package main

type User struct {
	Name      string
	Email     string
	AccountId string
}

type Repository struct {
	BitbucketID string
	Name        string
	FullName    string
	Private     bool
}

type Project struct {
	BitbucketID string
	Name        string
	Key         string
}

type Commit struct {
	Hash        string
	Date        string
	Message     string
	AuthorEmail string
	AuthorName  string
}

type BitbucketMappedEvent struct {
	UserDocuments      []User
	RepositoryDocument Repository
	ProjectDocument    Project
	CommitDocuments    []Commit
}
