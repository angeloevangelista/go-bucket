package util

import "encoding/json"

func DeserializeObject[T any](rawJson string) (*T, error) {
	var serializedObject T

	err := json.Unmarshal([]byte(rawJson), &serializedObject)

	if err != nil {
		return nil, err
	}

	return &serializedObject, nil
}
