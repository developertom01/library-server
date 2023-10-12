package scalers

import (
	"fmt"
	"io"
	"regexp"
)

type UUID string

func (uuid *UUID) UnmarshalGQL(v interface{}) error {
	// Define a regex pattern for UUID (version 4, simple format)
	regex := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)
	isMatch := regex.MatchString(string(*uuid))
	if !isMatch {
		return fmt.Errorf("Must be a valid UUID")
	}
	return nil
}

func (uuid UUID) MarshalGQL(w io.Writer) {
	w.Write([]byte(uuid))
}
