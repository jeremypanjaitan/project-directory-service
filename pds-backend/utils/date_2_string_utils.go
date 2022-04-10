package utils

import "time"

func Date2String(date time.Time) string {
	result := date.String()
	result = date.Format("2006-01-02 15:04:05")
	return result
}