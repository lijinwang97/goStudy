package utils

import (
	log "code.inke.cn/BackendPlatform/golang/logging"
	"time"
)

var (
	Ymd string
)

func StartAndEndTime(minute int) (string,string) {
	now :=time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()-2, now.Minute()-3-minute, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	endTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()-2, now.Minute()-3, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	return startTime,endTime
}

func String2TimeStampe(startTime string) int64 {
	unixTime, err := time.Parse(startTime, "2006-01-02 15:04:05")
	if err != nil {
		log.Errorf("parse string time to TimeStamp fail! stringTime:%s",startTime)
		return 0
	}
	return unixTime.Unix()
}
