package times

import "time"

func Format(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
func FormatYMD(t time.Time) string {
	return t.Format("2006-01-02")
}
func FormatByMill(t int64) string {
	return time.UnixMilli(t).Format("2006-01-02 15:04:05")
}

func ParseTime(timeStr string) int64 {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	return t.UnixMilli()
}
