package timex

import (
	"sync"
	"time"
)

var location = "Asia/Shanghai"
var setLocationOnce sync.Once
var local, _ = time.LoadLocation(location)

func SetLocation(loc string) {
	setLocationOnce.Do(func() {
		location = loc
		var err error
		local, err = time.LoadLocation(loc)
		if err != nil {
			panic(err)
		}
	})
}

func GetLocation() *time.Location {
	if local == nil {
		panic("location is not set")
	}
	return local
}

func Now() time.Time {
	return time.Now().In(GetLocation())
}

func Format(t time.Time) string {
	return t.Format(time.DateTime)
}
