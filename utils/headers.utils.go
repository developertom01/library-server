package utils

import (
	"fmt"
	"strings"
)

func ExtractBearerToken(token string) (string, error) {
	if !strings.Contains(token, "Bearer") {
		return "", fmt.Errorf("String provided is not a valid token")
	}

	return strings.Replace(token, "Bearer ", "", 1), nil
}
