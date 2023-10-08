package utils

import "time"

func ConvertTimeToIso(date time.Time) *string {
	timeString:=date.Format(time.RFC3339)
	return &timeString
}
