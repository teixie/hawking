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

// 时间格式化，支持"Y-m-d H:i:s"、"YYYY-mm-dd HH:ii:ss"等形式
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

// 获得Golang的time.Time类型的时间
func (h Hawking) Time() time.Time {
	return h.t
}

// 获得时间戳
func (h Hawking) Unix() int64 {
	return h.t.Unix()
}

// 增加时间
func (h Hawking) Add(d time.Duration) Hawking {
	h.t = h.t.Add(d)
	return h
}

// 设置时区
func SetLocation(loc *time.Location) {
	local = loc
}

// 获得时区
func GetLocation() *time.Location {
	if local != nil {
		return local
	}
	return time.Local
}

// 获得当前时间
func Now() Hawking {
	return Hawking{time.Now().In(GetLocation())}
}

// 获得今天的开始时间
func Today() Hawking {
	timeStr := time.Now().In(GetLocation()).Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, GetLocation())
	return Hawking{t}
}

// 获得明天的开始时间
func Tomorrow() Hawking {
	timeStr := time.Now().In(GetLocation()).Add(24 * time.Hour).Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, GetLocation())
	return Hawking{t}
}

// 获得昨天的开始时间
func Yesterday() Hawking {
	timeStr := time.Now().In(GetLocation()).Add(-24 * time.Hour).Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, GetLocation())
	return Hawking{t}
}
