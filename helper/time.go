package helper

import (
	"strconv"
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

func GetExpiredDateTime(t string) string {
	now := time.Now().UTC()
	yyyy, m, dd := now.Date()

	monthInt := int(m)
	monthStr := strconv.Itoa(monthInt)
	if monthInt < 10 {
		monthStr = "0" + strconv.Itoa(monthInt)
	}

	day := strconv.Itoa(dd)
	if dd < 10 {
		day = "0" + strconv.Itoa(dd)
	}

	if now.Hour() >= 15 {
		day = strconv.Itoa(dd + 1)
	}

	expiredDatetime := strconv.Itoa(yyyy) + "-" + monthStr + "-" + day + " " + t
	return expiredDatetime
}
