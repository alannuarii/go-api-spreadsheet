package utils

import (
	"strconv"
	"strings"
)

func StringToDecimal(str string) float64 {

	str = strings.Replace(str, ",", ".", -1)

	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}

	return floatValue
}
