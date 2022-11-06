package util

import "regexp"

func SanitizeBitbucketUUID(id string) string {
	return regexp.MustCompile(`{|}`).ReplaceAllString(id, "")
}
