package scalers

import (
	"fmt"
	"io"
	"regexp"
)

type Url string

func (url *Url) UnmarshalGQL(v interface{}) error {
	regex := regexp.MustCompile(`^(http|https):\/\/[^\s/$.?#].[^\s]*$`)
	isMatch := regex.MatchString(string(*url))
	if !isMatch {
		return fmt.Errorf("Must be a valid url")
	}
	return nil
}

func (url Url) MarshalGQL(w io.Writer) {
	w.Write([]byte(url))
}
