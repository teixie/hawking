package hawking

import (
	"regexp"
	"time"
)

var (
	local *time.Location
)

type Hawking struct {
	t time.Time
}

func (h Hawking) Format(fmtStr string) string {
	timeStr := h.t.String()
	o := map[string]string{
		"Y+": timeStr[0:4],
		"y+": timeStr[2:4],
		"m+": timeStr[5:7],
		"d+": timeStr[8:10],
		"H+": timeStr[11:13],
		"i+": timeStr[14:16],
		"s+": timeStr[17:19],
	}
	for k, v := range o {
		re, _ := regexp.Compile(k)
		fmtStr = re.ReplaceAllString(fmtStr, v)
	}
	return fmtStr
}

func (h Hawking) Time() time.Time {
	return h.t
}

// 设置时区
func SetLocation(loc *time.Location) {
	local = loc
}

// 获取时区
func GetLocation() *time.Location {
	if local != nil {
		return local
	}
	return time.Local
}

// 获取当前时间
func Now() Hawking {
	return Hawking{time.Now().In(GetLocation())}
}

// 获取今天的起始时间
func Today() Hawking {
	timeStr := Now().Time().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, GetLocation())
	return Hawking{t}
}

// 获取明天的起始时间
func Tomorrow() Hawking {
	timeStr := Now().Time().Add(24 * time.Hour).Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, GetLocation())
	return Hawking{t}
}

// 获取昨天的起始时间
func Yesterday() Hawking {
	timeStr := Now().Time().Add(-24 * time.Hour).Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, GetLocation())
	return Hawking{t}
}
