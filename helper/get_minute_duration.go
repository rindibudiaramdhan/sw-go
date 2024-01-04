package helper

import (
	"time"

	"github.com/rindibudiaramdhan/sw-go/helper/convert"
)

func GetMinuteDuration(expiredAtUtc time.Time) (int, error) {
	duration := time.Since(expiredAtUtc)
	minuteDuration := duration.Minutes() * -1
	minuteDurationInt, err := convert.FloatToInt(minuteDuration)
	if err != nil {
		return minuteDurationInt, err
	}
	return minuteDurationInt, nil
}
