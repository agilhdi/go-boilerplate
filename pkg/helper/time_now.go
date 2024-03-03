package helper

import "time"

func TimeToUnix() int64 {
	return time.Now().Unix()
}

func UnixToTime(unix int64) time.Time {
	return time.Unix(unix, 0)
}
