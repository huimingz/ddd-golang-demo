package datetime

import (
	"time"
)

var timeLocationGTM8 *time.Location

// Now get the current time in East 8.
func Now() time.Time {
	return time.Now().In(timeLocationGTM8)
}

func init() {
	// load time location of East 8 from tzData string.
	var err error
	timeLocationGTM8, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic("load ShanghaiTimeLocation failed from TZData, error: " + err.Error())
	}
}
