package utils

import "strings"

func ExtractBearerToken(token string) string {
	return strings.Split("Bearer ", token)[1]
}
