package util

import (
	"encoding/json"
)

func SerializeObject(objectToSerialize any) (*string, error) {
	jsonBytes, err := json.MarshalIndent(objectToSerialize, "", "\t")

	if err != nil {
		return nil, err
	}

	rawJson := string(jsonBytes)

	return &rawJson, nil
}
