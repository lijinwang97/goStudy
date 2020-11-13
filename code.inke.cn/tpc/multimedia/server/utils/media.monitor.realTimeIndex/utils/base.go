package utils

import (
	"fmt"
	"strconv"
)

func FloatRetain2(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func Int2String(value int) string {
	return strconv.Itoa(value)
}

