package utils

import (
	"fmt"
	"strconv"

	"github.com/developertom01/library-server/app/graphql/scalers"
	"github.com/google/uuid"
)

func ParseStringToUint(t string) (*uint, error) {
	parsedT, err := strconv.ParseUint(t, 10, 64)
	if err != nil {
		return nil, err
	}
	uintT := uint(parsedT)
	return &uintT, nil

}
func ParseMultipleStringToUint(t []string) ([]uint, error) {
	var v []uint
	for _, c := range t {
		a, err := ParseStringToUint(c)
		if err != nil {
			return []uint{}, nil
		}
		v = append(v, *a)

	}
	return v, nil

}

func ParseUintToString(t uint) *string {
	ts := fmt.Sprint(t)
	return &ts
}

func ParseScalerUuidToNativeUuid(id scalers.UUID) (uuid.UUID, error) {
	return uuid.Parse(string(id))
}
