package hawking

import (
	"regexp"
	"time"
)

const (
	timeLayoutYmdHis = "2006-01-02 15:04:05"
	timeLayoutYmd    = "2006-01-02"
)

var (
	local *time.Location
)

type Hawking struct {
	t time.Time
}

// 时间格式化，支持"Y-m-d H:i:s"、"YYYY-mm-dd HH:ii:ss"等形式，当不包含任何"YymdHis"字符时将使用原生方式
func (h Hawking) Format(fmtStr string) string {
	exists, err := regexp.Match("[YymdHis]+", []byte(fmtStr))
	if err == nil && !exists {
		return h.t.Format(fmtStr)
	}

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

// 时间是否为0
func (h Hawking) IsZero() bool {
	return h.t.IsZero()
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
	timeStr := time.Now().In(GetLocation()).Format(timeLayoutYmd)
	t, _ := time.ParseInLocation(timeLayoutYmd, timeStr, GetLocation())
	return Hawking{t}
}

// 获得明天的开始时间
func Tomorrow() Hawking {
	timeStr := time.Now().In(GetLocation()).Add(24 * time.Hour).Format(timeLayoutYmd)
	t, _ := time.ParseInLocation(timeLayoutYmd, timeStr, GetLocation())
	return Hawking{t}
}

// 获得昨天的开始时间
func Yesterday() Hawking {
	timeStr := time.Now().In(GetLocation()).Add(-24 * time.Hour).Format(timeLayoutYmd)
	t, _ := time.ParseInLocation(timeLayoutYmd, timeStr, GetLocation())
	return Hawking{t}
}

// 构建时间，支持Hawking/time.Time/时间字符串/时间戳
func Make(t interface{}) Hawking {
	if t == nil {
		return Hawking{}
	}

	switch t.(type) {
	case Hawking:
		return t.(Hawking)
	case time.Time:
		return Hawking{t.(time.Time)}
	case string:
		r, err := time.ParseInLocation(timeLayoutYmdHis, t.(string), GetLocation())
		if err == nil {
			return Hawking{r}
		}
	case int:
		return Hawking{time.Unix(int64(t.(int)), 0)}
	case int64:
		return Hawking{time.Unix(t.(int64), 0)}
	}

	return Hawking{}
}
