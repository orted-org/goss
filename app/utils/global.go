package utils

import "time"

func GetTimeFor(seconds int) time.Duration {
	return time.Second * time.Duration(seconds)
}
