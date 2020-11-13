package utils

import (
	log "code.inke.cn/BackendPlatform/golang/logging"
	"time"
)

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
)

var (
	Ymd string
)

func StartAndEndTime(minute int) (string, string) {
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()-2, now.Minute(), -1, 0, time.Local).Format("2006-01-02 15:04:05")
	//startTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()-2, now.Minute()-3-minute, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	endTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()-2, now.Minute(), 0, 0, time.Local).Format("2006-01-02 15:04:05")
	return startTime, endTime
}

func StartAndEndTimestamp(minute int) (int64, int64) {
	//now := time.Now()
	//startTime := time.Date(now.Year(), now.Month(), now.Day()-1, now.Hour(), now.Minute()-2*minute, -1, 0, time.Local).UnixNano() / 1e6
	//endTime := time.Date(now.Year(), now.Month(), now.Day()-1, now.Hour(), now.Minute(), 0, 0, time.Local).UnixNano() / 1e6
	var startTime int64
	var endTime int64
	startTime = 1605095619473
	endTime = 1605095719473
	return startTime, endTime
}

func String2TimeStampe(startTime string) int64 {
	unixTime, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	if err != nil {
		log.Errorf("parse string time to TimeStamp fail! stringTime:%s", startTime)
		return 0
	}
	return unixTime.Unix()
}

//时间戳转格式化时间
func GetUnix2Format(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(TIME_FORMAT)
}

//格式化时间转时间戳
func GetFormat2Unix(formatTime string) int64 {
	unixTime, _ := time.ParseInLocation(TIME_FORMAT, formatTime, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	return unixTime.Unix()
}

//获取指定时间
func GetAppointTime() string {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.Local).Format(TIME_FORMAT) //参数时间可修改
}

//获取当前格式化时间
func GetCurTimeFormat() string {
	return time.Now().Format(TIME_FORMAT)
}

//获取当前时间戳
func GetCurTimestamp() int64 {
	return time.Now().Unix()
}
