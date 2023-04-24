package fn

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	u := uuid.New()
	return strings.ReplaceAll(u.String(), "-", "")
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
