package utils

import "strings"

func IsUniqueConstraintViolated(err error) bool {
	if strings.Contains(err.Error(), "UNIQUE constraint failed") {
		return true
	}
	return false
}
