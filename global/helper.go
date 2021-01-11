package global

import (
	"math/rand"
	"time"
)

//这个包放置公共常量
const (
	DateTimeFormatStr = "2006-01-02 15:04:05"
	DateFormatStr = "2006-01-02"

)


func Hour2Unix(hour string) (time.Time, error) {
	return time.ParseInLocation(DateTimeFormatStr, time.Now().Format(DateFormatStr) + " " + hour, time.Local)
}

func UnixMilli() int64 {
	return time.Now().UnixNano() / 1e6
}
func GenerateRangeNum(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int63n(max - min) + min
	return randNum
}