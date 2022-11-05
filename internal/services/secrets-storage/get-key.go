package secrets_storage_service

import (
	"errors"
	"os"
)

func GetSecret(key string) (*string, error) {
	var storage = map[string]string{
		"BITBUCKET_LEGACY_CLIENT_ID":     os.Getenv("BITBUCKET_LEGACY_CLIENT_ID"),
		"BITBUCKET_LEGACY_CLIENT_SECRET": os.Getenv("BITBUCKET_LEGACY_CLIENT_SECRET"),
	}

	value := storage[key]

	if value == "" {
		return nil, errors.New("key not found")
	}

	return &value, nil
}
