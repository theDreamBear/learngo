package util

import (
	"fmt"
	"strconv"
	"time"
)

func TimeStampToYearMonthDay(stime int64) int32 {
	tm := time.Unix(stime, 0)

	ymdstring := fmt.Sprintf("%4d%02d%02d", tm.Year(), tm.Month(), tm.Day())
	if ymd, err := strconv.ParseInt(ymdstring, 10, 32); err == nil {
		return int32(ymd)
	} else {
		return -1
	}
}
