package dateutils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

// GetNow returns current time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns current date as string
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetDbString returns current date as string
func GetDbString() string {
	return GetNow().Format(apiDbLayout)
}
