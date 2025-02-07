// Package date_utils is a helper for manipulate date with time zone associated
package date_utils

import "time"

const apiDateLayout = "2006-01-02T15:04:05z"

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
