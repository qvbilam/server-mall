package utils

import "time"

const TemplateDateDefault = "YYYY-MM-DD HH:mm:ss"

func TimestampToCustomizeTimeDate(t int64, temp string) string {
	return time.Unix(t, 0).Local().Format(temp)
}

func TimestampToTimeDate(t int64) string {
	return TimestampToCustomizeTimeDate(t, TemplateDateDefault)
}

func TimeDateToTimestamp(d string) int64 {
	t, _ := time.Parse(TemplateDateDefault, d)
	return t.Unix()
}
