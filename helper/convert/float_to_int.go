package convert

import (
	"fmt"
	"strconv"
)

func FloatToInt(f float64) (int, error) {
	t := int(f)
	s := fmt.Sprintf("%.0f", f)
	i, err := strconv.Atoi(s)
	if err != nil {
		return t, err
	}
	return i, nil
}
