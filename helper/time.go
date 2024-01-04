package helper

import (
	"time"
)

/**
* layout is ... .
* value is ... .
* locationName is Zone Info
 */
func ParseInLocation(layout string, value string, locationName string) (*time.Time, error) {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return nil, err
	}

	expiredAt, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return nil, err
	}

	return &expiredAt, nil
}
